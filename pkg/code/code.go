package code

import "github.com/yongcycchen/mall-api/common/errcode"

const (
	SUCCESS                   = 200
	InvalidParams             = 400
	ERROR                     = 500
	IdNotEmpty                = 4001
	ErrorTokenEmpty           = 4002
	ErrorTokenInvalid         = 4003
	ErrorTokenExpire          = 4004
	ErrorUserNotExist         = 4005
	ErrorUserExist            = 4006
	ErrorUserPwd              = 4007
	ErrorMerchantNotExist     = 4008
	ErrorMerchantExist        = 4009
	ErrorShopBusinessExist    = 4010
	ErrorShopBusinessNotExist = 4011
	ErrorSkuCodeExist         = 4012
	ErrorSkuCodeNotExist      = 4013
	ErrorShopIdNotExist       = 4014
	ErrorShopIdExist          = 4015
	ErrorInviteCodeNotExist   = 4016
	DbDuplicateEntry          = 50000
	ErrorEmailSend            = 50001
	ErrorVerifyCodeEmpty      = 50002
	ErrorVerifyCodeInvalid    = 50003
	ErrorVerifyCodeExpire     = 50004
	ErrorSkuAmountNotEnough   = 50005
	ErrorVerifyCodeInterval   = 50006
	ErrorVerifyCodeLimited    = 50007
	UserBalanceNotEnough      = 600001
	UserAccountStateLock      = 600002
	UserAccountNotExist       = 600003
	MerchantAccountNotExist   = 600004
	MerchantAccountStateLock  = 600005
	DecimalParseErr           = 600000
	TransactionFailed         = 600010
	TxCodeNotExist            = 600011
	TradePayRun               = 600012
	TradePaySuccess           = 600013
	LogisticsRecordExist      = 600014
	LogisticsRecordNotExist   = 600015
	UserLoginNotAllow         = 600016
	TradePayExpire            = 600017
	TradeOrderTxCodeEmpty     = 600018
	TradeOrderExist           = 600019
	UserSettingInfoExist      = 600020
	UserSettingInfoNotExist   = 600021
	UserDeliveryInfoNotExist  = 600022
	TradeOrderNotMatchUser    = 600023
	SkuPriceVersionNotExist   = 600024
	OrderStateInvalid         = 600025
	OrderStateLock            = 600026
	OrderExpire               = 600027
	OrderPayCompleted         = 600028
	UserAccountStateInvalid   = 600029
	CommentsExist             = 600030
	CommentsNotExist          = 600031
	CommentsTagExist          = 600032
	CommentsTagNotExist       = 600033
	UserOrderNotExist         = 600034
	OutTradeEmpty             = 600035
)

const (
	Success               = 29000000
	ErrorServer           = 29000001
	UserNotExist          = 29000005
	UserExist             = 29000006
	DBDuplicateEntry      = 29000007
	MerchantExist         = 29000008
	MerchantNotExist      = 29000009
	AccountExist          = 29000010
	AccountNotExist       = 29000011
	UserPwdNotMatch       = 29000012
	UserDeliveryInfoExist = 29000013
	// UserDeliveryInfoNotExist = 29000014
	// TransactionFailed        = 29000015
	AccountStateLock       = 29000016
	AccountStateInvalid    = 29000017
	UserChargeRun          = 29000018
	UserChargeSuccess      = 29000019
	UserChargeTradeNoEmpty = 29000020
)

var ErrMap = make(map[int]string)

func init() {
	dict := map[int]string{
		Success:                  "OK",
		ErrorServer:              "服务器错误",
		UserNotExist:             "用户不存在",
		DBDuplicateEntry:         "Duplicate entry",
		UserExist:                "已存在用户记录，请勿重复创建",
		MerchantExist:            "商户认证材料已存在",
		MerchantNotExist:         "商户未提交材料",
		AccountExist:             "账户已存在",
		AccountNotExist:          "账户不存在",
		UserPwdNotMatch:          "用户密码不匹配",
		UserDeliveryInfoExist:    "用户物流交付信息存在",
		UserDeliveryInfoNotExist: "用户物流交付信息不存在",
		TransactionFailed:        "事务执行失败",
		AccountStateLock:         "用户账户锁定中",
		AccountStateInvalid:      "用户账户无效",
		UserChargeRun:            "本次充值交易正在进行中",
		UserChargeSuccess:        "本次充值交易已成功",
		UserChargeTradeNoEmpty:   "本次充值交易号为空",
	}
	errcode.RegisterErrMsgDict(dict)
	for key, _ := range dict {
		ErrMap[key] = dict[key]
	}
}

// func GetMsg(code int) string {
// 	v, ok := ErrMap[code]
// 	if !ok {
// 		return ""
// 	}
// 	return v
// }
