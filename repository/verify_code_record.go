package repository

import (
	"github.com/yongcycchen/mall-api/model/mysql"
	"github.com/yongcycchen/mall-api/vars"
)

func CreateVerifyCodeRecord(record *mysql.VerifyCodeRecord) (err error) {
	_, err = vars.DBEngineXORM.Table(mysql.TableVerifyCodeRecord).Insert(record)
	if err !=nil {
		return err
	}
	return
}
