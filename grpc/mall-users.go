package main

import (
	mallusers "github.com/yongcycchen/mall-api/startup/grpc/mall-users"
	"github.com/yongcycchen/mall-api/vars"
)

const APP_NAME = "mall-users"

func main() {
	application := &vars.GRPCApplication{
		Application: &vars.Application{
			Name:       APP_NAME,
			LoadConfig: mallusers.LoadConfig,
			SetupVars:  mallusers.SetupVars,
		},
		RegisterGRPCServer: mallusers.RegisterGRPCServer,
		RegisterGateway:    mallusers.RegisterGateway,
		RegisterHttpRoute:  mallusers.RegisterHttpRoute,
	}
	vars.MallusersApp = application
}
