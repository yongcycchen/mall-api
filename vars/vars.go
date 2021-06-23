package vars

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/yongcycchen/mall-api/common/log"
	"github.com/yongcycchen/mall-api/config/setting"
	"github.com/yongcycchen/mall-api/g2cache"
	"github.com/yongcycchen/mall-api/pkg/util/goroutine"
	"xorm.io/xorm"
)

var (
	App                   *WEBApplication
	DBEngineXORM          xorm.EngineInterface
	DBEngineGORM          *gorm.DB
	LoggerSetting         *setting.LoggerSettingS
	AccessLogger          log.LoggerContextIface
	ErrorLogger           log.LoggerContextIface
	BusinessLogger        log.LoggerContextIface
	ServerSetting         *setting.ServerSettingS
	JwtSetting            *setting.JwtSettingS
	MysqlSettingMicroMall *setting.MysqlSettingS
	RedisSettingMicroMall *setting.RedisSettingS
	G2CacheSetting        *setting.G2CacheSettingS
	EmailConfigSetting    *EmailConfigSettingS
	VerifyCodeSetting     *VerifyCodeSettingS
	RedisPoolMicroMall    *redis.Pool
	GPool                 *goroutine.Pool
	G2CacheEngine         *g2cache.G2Cache
)
