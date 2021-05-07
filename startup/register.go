package startup

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/router"
	"github.com/yongcycchen/mall-api/vars"
)

func RegisterHttpRoute() *gin.Engine {
	accessInfoLogger := &AccessInfoLogger{}
	accessErrLogger := &AccessErrLogger{}
	ginRouter := router.InitRouter(accessInfoLogger, accessErrLogger)
	return ginRouter
}

type AccessInfoLogger struct{}

func (a *AccessInfoLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Infof(context.Background(), "[gin-info] %s", p)
	return 0, nil
}

type AccessErrLogger struct{}

func (a *AccessErrLogger) Write(p []byte) (n int, err error) {
	vars.AccessLogger.Errorf(context.Background(), "[gin-err] %s", p)
	return 0, nil
}

// regist tasks
func RegisterTasks() []vars.CronTask {
	var tasks = make([]vars.CronTask, 0)
	tasks = append(tasks) // TestCronTask()
	return tasks
}
