package startup

import (
	"log"

	"github.com/yongcycchen/mall-api/config/setting"
	"github.com/yongcycchen/mall-api/config"
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
	config.MapConfig(SectionMysqlMicroMall,vars.MysqlSettingMicroMall)
	return nil
}
