package grpc

import (
	"github.com/yongcycchen/mall-api/vars"
)

const APP_NAME = "mall-users"

func main() {
	application := &vars.GRPCApplication{
		Application: &vars.Application{
			LoadConfig: startup.LoadConfig,
		},
	}
}
