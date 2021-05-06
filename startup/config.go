package startup

import (
	"log"

	"github.com/yongcycchen/mall-api/config"
	"github.com/yongcycchen/mall-api/config/setting"
	"github.com/yongcycchen/mall-api/vars"
)

const (
	SectionMysqlMicroMall = "mall-mysql"
	SectionRedisMicroMall = "mall-redis"
	SectionEmailConfig    = "email-config"
	SectionVerifyCode     = "mall-verify_code"
	SectionG2Cache        = "mall-g2cache"
)

func LoadConfig() error {

	// MySQL
	log.Printf("[info] Load default config %s", SectionMysqlMicroMall)
	vars.MysqlSettingMicroMall = new(setting.MysqlSettingS)
	config.MapConfig(SectionMysqlMicroMall, vars.MysqlSettingMicroMall)

	// Redis
	log.Printf("[info] Load defalut config %s", SectionRedisMicroMall)
	vars.RedisSettingMicroMall = new(setting.RedisSettingS)
	config.MapConfig(SectionRedisMicroMall, vars.RedisSettingMicroMall)

	//G2Cache
	log.Printf("[info] Load default config %s", SectionG2Cache)
	vars.G2CacheSetting = new(setting.G2CacheSettingS)
	config.MapConfig(SectionG2Cache, vars.G2CacheSetting)

	//email
	log.Printf("[info] Load default config %s", SectionEmailConfig)
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)

	//Verification code
	log.Printf("[info] Load default config %s", SectionVerifyCode)
	vars.VerifyCodeSetting = new(vars.VerifyCodeSettingS)
	config.MapConfig(SectionVerifyCode, vars.VerifyCodeSetting)
	return nil
}
