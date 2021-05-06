package startup

import (
	"gitee.com/kelvins-io/common/log"
	"github.com/yongcycchen/mall-api/internal/setup"
	"github.com/yongcycchen/mall-api/pkg/util/goroutine"
	"github.com/yongcycchen/mall-api/vars"
)

func SetupVars() error {
	var err error

	vars.ErrorLogger, err = log.GetErrLogger("err")
	if err != nil {
		return err
	}

	vars.BusinessLogger, err = log.GetBusinessLogger("business")
	if err != nil {
		return err
	}

	vars.AccessLogger, err = log.GetAccessLogger("access")
	if err != nil {
		return err
	}

	if vars.MysqlSettingMicroMall != nil && vars.MysqlSettingMicroMall.Host != "" {
		vars.DBEngineXORM, err = setup.NewMySQLXORMEngine(vars.MysqlSettingMicroMall)
		if err != nil {
			return err
		}
		vars.DBEngineGORM, err = setup.NewMySQLGORMEngine(vars.MysqlSettingMicroMall)
		if err != nil {
			return err
		}
	}

	if vars.RedisSettingMicroMall != nil && vars.RedisSettingMicroMall.Host != "" {
		vars.RedisPoolMicroMall, err = setup.NewRedis(vars.RedisSettingMicroMall)
		if err != nil {
			return err
		}
	}

	if vars.G2CacheSetting != nil && vars.G2CacheSetting.RedisConfDSN != "" {
		vars.G2CacheEngine, err = setup.NewG2Cache(vars.G2CacheSetting, nil, nil)
		if err != nil {
			return err
		}
	}
	vars.GPool = goroutine.NewPool(20, 1000)
	return nil
}