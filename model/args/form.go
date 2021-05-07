package args

type GenVerifyCodeArgs struct {
	CountryCode  string `form:"country_code" json:"country_code"`
	Phone        string `form:"phone" json:"phone"`
	BusinessType int    `form:"business_type" json:"business_type"`
	ReceiveEmail string `form:"receive_email" json:"receive_email"`
}
