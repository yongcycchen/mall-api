package main

import (
	"github.com/yongcycchen/mall-api/app"
	"github.com/yongcycchen/mall-api/startup"
	"github.com/yongcycchen/mall-api/vars"
)

const AppName = "mall-api"

func main() {
	application := &vars.WEBApplication{
		Application: &vars.Application{
			Name:       AppName,
			LoadConfig: startup.LoadConfig,
			SetupVars:  startup.SetupVars,
			StopFunc:   startup.SetStopFunc,
		},
		RegisterHttpRoute: startup.RegisterHttpRoute,
		RegisterTasks:     startup.RegisterTasks,
	}
	app.RunApplication(application)
}
