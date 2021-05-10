package args

const (
	Unknown            = 0
	VerifyCodeRegister = 1
	VerifyCodeLogin    = 2
	VerifyCodePassword = 3
	VerifyCodeTemplate = "【%v】verify code %v，within %v，%v minutes"
)

const (
	UserStateEventTypeRegister  = 10010
	UserStateEventTypeLogin     = 10011
	UserStateEventTypeLogout    = 10012
	UserStateEventTypePwdModify = 10013
)

const (
	RpcServiceMallUsers     = "mall-users"     // users service
	RpcServiceMallShop      = "mall-shop"      // shop service
	RpcServiceMallSku       = "mall-sku"       // sku service
	RpcServiceMallTrolley   = "mall-trolley"   // trolley service
	RpcServiceMallOrder     = "mall-order"     // order service
	RpcServiceMallPay       = "mall-pay"       // pay service
	RpcServiceMallLogistics = "mall-logistics" // logistics service
	RpcServiceMallComments  = "mall-comments"  // comments service
)

const (
	CNY = 0
	USD = 1
)

var (
	VerifyCodeTypes = []int{VerifyCodeRegister, VerifyCodeLogin, VerifyCodePassword}
	CoinTypes       = []int{CNY, USD}
)

var MsgFlags = map[int]string{
	Unknown:                     "Unknown",
	VerifyCodeRegister:          "VerifyCodeRegister",
	VerifyCodeLogin:             "VerifyCodeLogin",
	VerifyCodePassword:          "VerifyCodePassword:",
	UserStateEventTypeRegister:  "Register",
	UserStateEventTypePwdModify: "Password Modify",
	UserStateEventTypeLogin:     "Login",
	UserStateEventTypeLogout:    "Logout",
}

type CommonBusinessMsg struct {
	Type int    `json:"type"`
	Tag  string `json:"tag"`
	UUID string `json:"uuid"`
	Msg  string `json:"msg"`
}

type UserRegisterNotice struct {
	CountryCode string `json:"country_code"`
	Phone       string `json:"phone"`
	Time        string `json:"time"`
	State       int    `json:"state"`
}

type UserStateNotice struct {
	Uid  int    `json:"uid"`
	Time string `json:"time"`
}

type UserOnlineState struct {
	Uid   int    `json:"uid"`
	State string `json:"state"`
	Time  string `json:"time"`
}

type SkuInventoryInfo struct {
	SkuCode       string `json:"sku_code"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	Title         string `json:"title"`
	SubTitle      string `json:"sub_title"`
	Desc          string `json:"desc"`
	Production    string `json:"production"`
	Supplier      string `json:"supplier"`
	Category      int32  `json:"category"`
	Color         string `json:"color"`
	ColorCode     int32  `json:"color_code"`
	Specification string `json:"specification"`
	DescLink      string `json:"desc_link"`
	State         int32  `json:"state"`
	Amount        int64  `json:"amount"`
	ShopId        int64  `json:"shop_id"`
	Version       int64  `json:"version"`
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}
