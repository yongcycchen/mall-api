package repository

import (
	"github.com/yongcycchen/mall-api/model/mysql"
	"github.com/yongcycchen/mall-api/vars"
)

func CreateVerifyCodeRecord(record *mysql.VerifyCodeRecord) (err error) {
	_, err = vars.DBEngineXORM.Table(mysql.TableVerifyCodeRecord).Insert(record)
	if err != nil {
		return err
	}
	return
}

func GetVerifyCode(businessType int, countryCode, phone, verifyCode string) (*mysql.VerifyCodeRecord, error) {
	var result mysql.VerifyCodeRecord
	var err error
	_, err = vars.DBEngineXORM.Table(mysql.TableVerifyCodeRecord).
		Select("id,expire").
		Where("business_type = ? AND country_code = ? AND phone = ? AND verify_code = ?", businessType, countryCode, phone, verifyCode).
		Desc("id").
		Get(&result)
	return &result, err
}
