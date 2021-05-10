package router

import (
	"io"
	"os"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/middleware"
	v1 "github.com/yongcycchen/mall-api/router/api/v1"
	"github.com/yongcycchen/mall-api/router/process"
	"github.com/yongcycchen/mall-api/vars"
)

func InitRouter(accessInfoLogger, accessErrLogger io.Writer) *gin.Engine {
	gin.DefaultWriter = io.MultiWriter(os.Stdout, accessInfoLogger)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, accessErrLogger)

	gin.SetMode(gin.ReleaseMode)
	if vars.ServerSetting != nil && vars.ServerSetting.Mode != "" {
		gin.SetMode(vars.ServerSetting.Mode)
	}

	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/", v1.IndexApi)
	pprof.Register(r, "/debug")
	r.GET("/debug/metrics", process.MetricsApi)
	r.GET("/ping", v1.PingApi)
	r.Static("/static", "./static")
	//???什么意思
	apiG := r.Group("/api")
	apiV1 := apiG.Group("/v1")
	apiV1.POST("/verify_code/send", v1.GetVerifyCodeApi)            // send virify code
	apiV1.POST("/register", v1.RegisterUserApi)                     // register
	apiV1.POST("/login/verify_code", v1.LoginUserWithVerifyCodeApi) // verify code login
	apiV1.POST("/login/pwd", v1.LoginUserWithPwdApi)                // password login
	apiUser := apiV1.Group("/user")
	apiUser.Use(middleware.CheckUserToken())
	{
	}
	return r
}
