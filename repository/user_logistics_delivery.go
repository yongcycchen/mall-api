package repository

import (
	"github.com/yongcycchen/mall-api/model/mysql"
	"github.com/yongcycchen/mall-api/vars"
	"xorm.io/xorm"
)

func CreateUserLogisticsDelivery(model *mysql.UserLogisticsDelivery) error {
	_, err := vars.XORM_DBEngine.Table(mysql.TableUserLogisticsDelivery).Insert(model)
	return err
}

func CreateUserLogisticsDeliveryByTx(tx *xorm.Session, model *mysql.UserLogisticsDelivery) error {
	_, err := tx.Table(mysql.TableUserLogisticsDelivery).Insert(model)
	return err
}

func CheckUserLogisticsDelivery(uid int64, ids []int32) ([]mysql.UserLogisticsDelivery, error) {
	var result = make([]mysql.UserLogisticsDelivery, 0)
	err := vars.XORM_DBEngine.Table(mysql.TableUserLogisticsDelivery).Select("id").Where("uid = ?", uid).In("id", ids).Find(&result)
	return result, err
}

func GetUserLogisticsDelivery(sqlSelect string, id int64) (*mysql.UserLogisticsDelivery, error) {
	var model mysql.UserLogisticsDelivery
	_, err := vars.XORM_DBEngine.Table(mysql.TableUserLogisticsDelivery).Select(sqlSelect).Where("id = ?", id).Get(&model)
	return &model, err
}

func GetUserLogisticsDeliveryList(sqlSelect string, uid int64) ([]mysql.UserLogisticsDelivery, error) {
	var result = make([]mysql.UserLogisticsDelivery, 0)
	session := vars.XORM_DBEngine.Table(mysql.TableUserLogisticsDelivery).Select(sqlSelect).Where("uid = ?", uid)
	err := session.Find(&result)
	return result, err
}

func UpdateUserLogisticsDelivery(where, maps interface{}) (int64, error) {
	return vars.XORM_DBEngine.Table(mysql.TableUserLogisticsDelivery).Where(where).Update(maps)
}

func UpdateUserLogisticsDeliveryByTx(tx *xorm.Session, where, maps interface{}) (int64, error) {
	return tx.Table(mysql.TableUserLogisticsDelivery).Where(where).Update(maps)
}
