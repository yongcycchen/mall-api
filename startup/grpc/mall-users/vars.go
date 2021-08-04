package mallusers

import (
	"github.com/yongcycchen/mall-api/pkg/util/goroutine"
	"github.com/yongcycchen/mall-api/setup"
	"github.com/yongcycchen/mall-api/vars"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	if vars.QueueAMQPSettingUserRegisterNotice != nil && vars.QueueAMQPSettingUserRegisterNotice.Broker != "" {
		vars.QueueServerUserRegisterNotice, err = setup.NewAMQPQueue(vars.QueueAMQPSettingUserRegisterNotice, nil)
		if err != nil {
			return err
		}
	}

	if vars.QueueAMQPSettingUserStateNotice != nil && vars.QueueAMQPSettingUserStateNotice.Broker != "" {
		vars.QueueServerUserStateNotice, err = setup.NewAMQPQueue(vars.QueueAMQPSettingUserStateNotice, nil)
		if err != nil {
			return err
		}
	}

	vars.GPool = goroutine.NewPool(20, 100)
	return nil
}
