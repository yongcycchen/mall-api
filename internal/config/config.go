package config

import (
	"flag"
	"log"

	"github.com/go-ini/ini"
	"github.com/yongcycchen/mall-api/config/setting"
	"github.com/yongcycchen/mall-api/vars"
)

const (
	// ConfFileName defines config file name.
	ConfFileName = "./etc/app.ini"
	// SectionServer is a section name for grpc server.
	SectionServer = "web-server"
	// SectionLogger is a section name for logger.
	SectionLogger = "web-logger"
	// jwt is token vaild
	SectionJwt = "web-jwt"
	// SectionMysql is a sectoin name for mysql.
	SectionMysql = "web-mysql"
	// SectionRedis is a section name for redis.
	SectionRedis = "web-redis"
)

// cfg reads file app.ini.
var (
	cfg      *ini.File
	flagConf = flag.String("conf_file", "", "Set app config.")
)

func LoadDefaultConfig(application *vars.Application) error {
	// Setup cfg object
	flag.Parse()
	var err error
	var confFile = ConfFileName
	if *flagConf != "" {
		confFile = *flagConf
	}
	cfg, err = ini.Load(confFile)
	if err != nil {
		return err
	}

	// Setup default settings
	for _, sectionName := range cfg.SectionStrings() {
		if sectionName == SectionServer {
			log.Printf("[info] Load default config %s", sectionName)
			vars.ServerSetting = new(setting.ServerSettingS)
			MapConfig(sectionName, vars.ServerSetting)
			continue
		}
		if sectionName == SectionLogger {
			log.Printf("[info] Load default config %s", sectionName)
			vars.LoggerSetting = new(setting.LoggerSettingS)
			MapConfig(sectionName, vars.LoggerSetting)
			continue
		}
		if sectionName == SectionJwt {
			log.Printf("[info] Load default config %s", sectionName)
			vars.JwtSetting = new(setting.JwtSettingS)
			MapConfig(sectionName, vars.JwtSetting)
		}
	}
	return nil
}

func MapConfig(section string, v interface{}) {
	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Fatalf("[err] Fail to parse '%s': %v", section, err)
	}
	err = sec.MapTo(v)
	if err != nil {
		log.Fatalf("[err] %s section map to setting err: %v", section, err)
	}
}
