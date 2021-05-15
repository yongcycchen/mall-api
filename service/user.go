package service

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/yongcycchen/mall-api/model/args"
	"github.com/yongcycchen/mall-api/model/mysql"
	"github.com/yongcycchen/mall-api/pkg/code"
	"github.com/yongcycchen/mall-api/pkg/util"
	"github.com/yongcycchen/mall-api/pkg/util/email"
	"github.com/yongcycchen/mall-api/proto/mall_users_proto/users"
	"github.com/yongcycchen/mall-api/repository"
	"github.com/yongcycchen/mall-api/vars"
)

func checkVerifyCodeLimit(limiter repository.CheckVerifyCodeLimiter, key string, limitCount int) (int, int) {
	if limitCount <= 0 {
		limitCount = repository.DefaultVerifyCodeSendPeriodLimitCount
	}
	count, err := limiter.GetVerifyCodePeriodLimitCount(key)
	if err != nil {
		return code.ERROR, count
	}
	if count >= limitCount {
		return code.ErrorVerifyCodeLimited, count
	}
	intervalTime, err := limiter.GetVerifyCodeInterval(key)
	if err != nil {
		return code.ERROR, count
	}
	if intervalTime == 0 {
		return code.SUCCESS, count
	}
	return code.ErrorVerifyCodeInterval, count
}

func GenVerifyCode(ctx context.Context, req *args.GenVerifyCodeArgs) (retCode int, verifyCode string) {
	retCode = code.SUCCESS
	var (
		err     error
		limiter = new(repository.CheckVerifyCodeMapCacheLimiter)
	)
	fmt.Printf("check 1\n")
	limitKey := fmt.Sprintf("%s%s", req.CountryCode, req.Phone)
	limitRetCode, limitCount := checkVerifyCodeLimit(limiter, limitKey, vars.VerifyCodeSetting.SendPeriodLimitCount)
	if limitRetCode != code.SUCCESS {
		vars.ErrorLogger.Infof(ctx, "checkVerifyCodeLimit %v %v is limited", req.CountryCode, req.Phone)
		retCode = limitRetCode
		return
	}
	fmt.Printf("check 2\n")
	serverName := args.RpcServiceMallUsers
	conn, err := util.GetGrpcClient(serverName)
	fmt.Printf("check 3\n")
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	fmt.Printf("check 4\n")
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	userReq := &users.GetUserInfoByPhoneRequest{
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
	}
	userRsp, err := client.GetUserInfoByPhone(ctx, userReq)
	if err != nil || userRsp.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
	verifyCode = strconv.Itoa(rand.Intn(1000000))

	verifyCodeRecord := mysql.VerifyCodeRecord{
		Uid:          int(userRsp.Info.Uid),
		BusinessType: req.BusinessType,
		VerifyCode:   verifyCode,
		Expire:       int(time.Now().Add(time.Duration(vars.VerifyCodeSetting.ExpireMinute) * time.Minute).Unix()),
		CountryCode:  req.CountryCode,
		Phone:        req.Phone,
		Email:        req.ReceiveEmail,
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	err = repository.CreateVerifyCodeRecord(&verifyCodeRecord)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "CreateVerifyCodeRecord err: %v, req: %+v", err, req)
		retCode = code.ERROR
		return
	}

	key := fmt.Sprintf("%s:%s-%s-%d", mysql.TableVerifyCodeRecord, req.CountryCode, req.Phone, req.BusinessType)
	err = vars.G2CacheEngine.Set(key, &verifyCodeRecord, 60*vars.VerifyCodeSetting.ExpireMinute, false)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "G2CacheEngine Set err: %v, key: %s,val: %+v", err, key, verifyCodeRecord)
		retCode = code.ERROR
		return
	}
	if req.ReceiveEmail != "" {
		job := func() {
			notice := fmt.Sprintf(args.VerifyCodeTemplate, vars.App.Name, verifyCode,
				args.GetMsg(req.BusinessType), vars.VerifyCodeSetting.ExpireMinute)
			err = email.SendEmailNotice(ctx, req.ReceiveEmail, vars.App.Name, notice)
			if err != nil {
				vars.ErrorLogger.Errorf(ctx, "SendEmailNotice err: %v, req: %+v", err, req)
				return
			}
		}
		vars.GPool.SendJob(job)
	}

	err = limiter.SetVerifyCodeInterval(limitKey, vars.VerifyCodeSetting.SendIntervalExpireSecond)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SetVerifyCodeInterval err: %v, req: %+v", err, req)
		retCode = code.ERROR
		return
	}
	err = limiter.SetVerifyCodePeriodLimitCount(limitKey, limitCount+1, vars.VerifyCodeSetting.SendPeriodLimitExpireSecond)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "SetVerifyCodePeriodLimitCount err: %v, req: %+v", err, req)
		retCode = code.ERROR
		return
	}
	return
}

func CreateUser(ctx context.Context, req *args.RegisterUserArgs) (*args.RegisterUserRsp, int) {
	var result args.RegisterUserRsp
	// check verify code
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeRegister,
		countryCode:  req.CountryCode,
		phone:        req.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return &result, retCode
	}
	serverName := args.RpcServiceMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return &result, code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	checkUserReq := users.CheckUserByPhoneRequest{
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
	}
	checkResult, err := client.CheckUserByPhone(ctx, &checkUserReq)
	if err != nil || checkResult.Common.Code != users.RetCode_SUCCESS {
		vars.ErrorLogger.Errorf(ctx, "CheckUserByPhone %v,err: %v,r : %+v",
			serverName, checkUserReq)
		return &result, code.ERROR
	}
	if checkResult.IsExist {
		return &result, code.ErrorUserExist
	}
	inviteId := int64(0)
	if req.InviteCode != "" {
		inviteUserReq := &users.GetUserByInviteCodeRequest{
			InviteCode: req.InviteCode,
		}
		inviteUser, err := client.GetUserInfoByInviteCode(ctx, inviteUserReq)
		if err != nil || inviteUser.Common.Code != users.RetCode_SUCCESS {
			vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode %v,err: %v,r : %+v", serverName, inviteUserReq)
			return &result, code.ERROR
		}
		if inviteUser.Info.Uid <= 0 {
			return &result, code.ErrorInviteCodeNotExist
		}
		inviteId = int64(int(inviteUser.Info.Uid))
	}
	// register req
	registerReq := &users.RegisterRequest{
		UserName:    req.UserName,
		Sex:         int32(req.Sex),
		CountryCode: req.CountryCode,
		Phone:       req.Phone,
		Email:       req.Email,
		IdCardNo:    req.IdCardNo,
		InviterUser: inviteId,
		ContactAddr: req.ContactAddr,
		Age:         int32(req.Age),
		Password:    req.Password,
	}
	registerRsp, err := client.Register(ctx, registerReq)
	if err != nil || registerRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "GetUserInfoByInviteCode %v,err: %v,r : %+v", serverName, registerReq)
		return &result, code.ERROR
	}
	switch registerRsp.Common.Code {
	case users.RetCode_USER_EXIST:
		return &result, code.ErrorUserExist
	}
	result.InviteCode = registerRsp.Result.InviteCode
	return &result, code.SUCCESS
}

type checkVerifyCodeArgs struct {
	businessType                   int
	countryCode, phone, verifyCode string
}

func checkVerifyCode(ctx context.Context, req *checkVerifyCodeArgs) int {
	key := fmt.Sprintf("%s:%s-%s-%d", mysql.TableVerifyCodeRecord, req.countryCode, req.phone, req.businessType)
	var obj mysql.VerifyCodeRecord
	err := vars.G2CacheEngine.Get(key, 60*vars.VerifyCodeSetting.ExpireMinute, &obj, func() (interface{}, error) {
		record, err := repository.GetVerifyCode(req.businessType, req.countryCode, req.phone, req.verifyCode)
		return record, err
	})
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetVerifyCode err: %v, req: %+v", err, req)
		return code.ERROR
	}

	if obj.Id == 0 {
		return code.ErrorVerifyCodeInvalid
	}
	if int64(obj.Expire) < time.Now().Unix() {
		return code.ErrorVerifyCodeExpire
	}
	return code.SUCCESS
}

func LoginUserWithVerifyCode(ctx context.Context, req *args.LoginUserWithVerifyCodeArgs) (string, int) {
	var token string
	reqCheckVerifyCode := checkVerifyCodeArgs{
		businessType: args.VerifyCodeLogin,
		countryCode:  req.CountryCode,
		phone:        req.Phone,
		verifyCode:   req.VerifyCode,
	}
	if retCode := checkVerifyCode(ctx, &reqCheckVerifyCode); retCode != code.SUCCESS {
		return token, retCode
	}
	serverName := args.RpcServiceMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	loginReq := &users.LoginUserRequest{
		LoginType: users.LoginType_VERIFY_CODE,
		LoginInfo: &users.LoginUserRequest_VerifyCode{
			VerifyCode: &users.LoginVerifyCode{
				Phone: &users.MobilePhone{
					CountryCode: req.CountryCode,
					Phone:       req.Phone,
				},
				VerifyCode: req.VerifyCode,
			},
		},
	}
	loginRsp, err := client.LoginUser(ctx, loginReq)
	if err != nil || loginRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "LoginUser %v,err: %v,r : %+v", serverName, loginReq)
		return "", code.ERROR
	}
	token = loginRsp.IdentityToken
	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ErrorUserNotExist
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ErrorUserPwd
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.UserLoginNotAllow
	}

	return token, code.SUCCESS
}

func LoginUserWithPwd(ctx context.Context, req *args.LoginUserWithPwdArgs) (string, int) {
	var token string
	serverName := args.RpcServiceMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return "", code.ERROR
	}
	defer conn.Close()
	client := users.NewUsersServiceClient(conn)
	loginReq := &users.LoginUserRequest{
		LoginType: users.LoginType_PWD,
		LoginInfo: &users.LoginUserRequest_Pwd{
			Pwd: &users.LoginByPassword{
				LoginKind: users.LoginPwdKind_MOBILE_PHONE,
				Info: &users.LoginByPassword_Phone{
					Phone: &users.MobilePhone{
						CountryCode: req.CountryCode,
						Phone:       req.Phone,
					},
				},
				Pwd: req.Password,
			},
		},
	}
	loginRsp, err := client.LoginUser(ctx, loginReq)
	if err != nil || loginRsp.Common.Code == users.RetCode_ERROR {
		vars.ErrorLogger.Errorf(ctx, "LoginUser %v,err: %v,r : %+v", serverName, loginReq)
		return "", code.ERROR
	}
	token = loginRsp.IdentityToken
	switch loginRsp.Common.Code {
	case users.RetCode_USER_NOT_EXIST:
		return "", code.ErrorUserNotExist
	case users.RetCode_USER_PWD_NOT_MATCH:
		return "", code.ErrorUserPwd
	case users.RetCode_USER_LOGIN_NOT_ALLOW:
		return "", code.UserLoginNotAllow
	}

	return token, code.SUCCESS
}
