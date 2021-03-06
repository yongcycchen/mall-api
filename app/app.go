package app

import (
	"fmt"
	"os"
	"time"

	"github.com/yongcycchen/mall-api/common/log"
	"github.com/yongcycchen/mall-api/internal/logging"
	"github.com/yongcycchen/mall-api/vars"
)

// server name

func initApplication(application *vars.Application) error {
	const DefaultLoggerRootPath = "./logs"
	const DefaultLoggerLevel = "debug"

	rootPath := DefaultLoggerRootPath
	if vars.LoggerSetting != nil && vars.LoggerSetting.RootPath != "" {
		rootPath = vars.LoggerSetting.RootPath
	}
	loggerLevel := DefaultLoggerLevel
	if vars.LoggerSetting != nil && vars.LoggerSetting.Level != "" {
		loggerLevel = vars.LoggerSetting.Level
	}
	err := log.InitGlobalConfig(rootPath, loggerLevel, application.Name)
	if err != nil {
		return fmt.Errorf("log.InitGlobalConfig: %v", err)
	}
	return nil
}

func appShutdown(application *vars.Application) error {
	if application.StopFunc != nil {
		return application.StopFunc()
	}
	return nil
}

func appPrepareForceExit() {
	time.AfterFunc(10*time.Second, func() {
		logging.Info("App server Shutdown timeout")
		os.Exit(1)
	})
}

// setup Common Vars
func setupCommonVars(application *vars.WEBApplication) error {
	if vars.ServerSetting != nil {
		vars.App.EndPort = vars.ServerSetting.EndPort
	}
	return nil
}
