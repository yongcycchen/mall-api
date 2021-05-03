package main

import (
	"github.com/yongcycchen/mall-api/startup"
	"github.com/yongcycchen/mall-api/vars"
)

const AppName = "mall-api"

func main() {
	application := &vars.WEBApplication{
		Application: &vars.Application{
			Name: AppName,
			LoadConfig: startup.LoadConfig,
		},
	}
}
