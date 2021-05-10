package args

import "github.com/astaxie/beego/validation"

type GenVerifyCodeArgs struct {
	CountryCode  string `form:"country_code" json:"country_code"`
	Phone        string `form:"phone" json:"phone"`
	BusinessType int    `form:"business_type" json:"business_type"`
	ReceiveEmail string `form:"receive_email" json:"receive_email"`
}

type RegisterUserArgs struct {
	UserName    string `form:"user_name" json:"user_name"`
	Password    string `form:"password" json:"password"`
	Sex         int    `form:"sex" json:"sex"`
	Email       string `form:"email" json:"email"`
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	Age         int    `json:"age" form:"age"`
	ContactAddr string `form:"contact_addr" json:"contact_addr"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
	IdCardNo    string `form:"id_card_no" json:"id_card_no"`
	InviteCode  string `form:"invite_code" json:"invite_code"`
}

type RegisterUserRsp struct {
	InviteCode string `json:"invite_code"` // RegisterUser invate_code
}

type LoginUserWithVerifyCodeArgs struct {
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	VerifyCode  string `form:"verify_code" json:"verify_code"`
}

func (t *LoginUserWithVerifyCodeArgs) Valid(v *validation.Validation) {
	if len(t.CountryCode) < 1 {
		v.SetError("CountryCode", "not less than 2 figure")
	}
	if len(t.Phone) < 10 {
		v.SetError("Phone", "not less than 10 figure")
	}
	if len(t.VerifyCode) < 6 {
		v.SetError("VerifyCode", "verify code not less than 6 figure")
	}
}

type LoginUserWithPwdArgs struct {
	CountryCode string `form:"country_code" json:"country_code"`
	Phone       string `form:"phone" json:"phone"`
	Password    string `form:"password" json:"password"`
}

func (t *LoginUserWithPwdArgs) Valid(v *validation.Validation) {
	if len(t.Password) < 6 {
		v.SetError("Password", "not less than 6 figure")
	}
	if len(t.CountryCode) < 1 {
		v.SetError("CountryCode", "not less than 2 figure")
	}
	if len(t.Phone) < 10 {
		v.SetError("Phone", "not less than 10 figure")
	}
}
