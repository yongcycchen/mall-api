package mallusers

import (
	"github.com/yongcycchen/mall-api/config"
	"github.com/yongcycchen/mall-api/config/setting"
	"github.com/yongcycchen/mall-api/vars"
)

const (
	SectionEmailConfig             = "email-config"
	SectionQueueUserRegisterNotice = "queue-user-register-notice"
	SectionQueueUserStateNotice    = "queue-user-state-notice"
	SectionJwt                     = "web-jwt"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 加载email数据源
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 用户注册通知
	vars.QueueAMQPSettingUserRegisterNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserRegisterNotice, vars.QueueAMQPSettingUserRegisterNotice)
	// 用户事件通知
	vars.QueueAMQPSettingUserStateNotice = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionQueueUserStateNotice, vars.QueueAMQPSettingUserStateNotice)
	// 用户认证token
	vars.JwtSetting = new(setting.JwtSettingS)
	config.MapConfig(SectionJwt, vars.JwtSetting)
	return nil
}
