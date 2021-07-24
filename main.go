package main

import (
	"github.com/yongcycchen/mall-api/app"
	"github.com/yongcycchen/mall-api/startup/web"
	"github.com/yongcycchen/mall-api/vars"
)

const AppName = "mall-api"

func main() {
	application := &vars.WEBApplication{
		Application: &vars.Application{
			Name:       AppName,
			LoadConfig: web.LoadConfig,
			SetupVars:  web.SetupVars,
			StopFunc:   web.SetStopFunc,
		},
		RegisterHttpRoute: web.RegisterHttpRoute,
		RegisterTasks:     web.RegisterTasks,
	}
	app.RunApplication(application)
}
