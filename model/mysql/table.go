package mysql

import (
	"time"
)

const (
	TableVerifyCodeRecord = "verify_code_record"
)

type VerifyCodeRecord struct {
	Id           int       `xorm:"'id' not null pk autoincr comment('new id') INT"`
	Uid          int       `xorm:"'uid' not null comment('user UID') INT"`
	BusinessType int       `xorm:"'business_type' comment('business type，1-register/login，2-buy ') TINYINT"`
	VerifyCode   string    `xorm:"'verify_code' comment('verify code') index CHAR(6)"`
	Expire       int       `xorm:"'expire' comment('expire time unix') INT"`
	CountryCode  string    `xorm:"'country_code' comment('country_code') index(country_code_phone_index) CHAR(5)"`
	Phone        string    `xorm:"'phone' comment('phone') index(country_code_phone_index) CHAR(11)"`
	Email        string    `xorm:"'email' comment('email') index VARCHAR(255)"`
	CreateTime   time.Time `xorm:"'create_time' comment('create time') DATETIME"`
	UpdateTime   time.Time `xorm:"'update_time' comment('update time') DATETIME"`
}
