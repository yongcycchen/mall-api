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
	limitKey := fmt.Sprintf("%s%s", req.CountryCode, req.Phone)
	limitRetCode, limitCount := checkVerifyCodeLimit(limiter, limitKey, vars.VerifyCodeSetting.SendPeriodLimitCount)
	if limitRetCode != code.SUCCESS {
		vars.ErrorLogger.Infof(ctx, "checkVerifyCodeLimit %v %v is limited", req.CountryCode, req.Phone)
		retCode = limitRetCode
		return
	}
	serverName := args.RpcServiceMallUsers
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		vars.ErrorLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		retCode = code.ERROR
		return
	}
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
}
