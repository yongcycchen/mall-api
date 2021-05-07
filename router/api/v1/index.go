package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/pkg/app"
	"github.com/yongcycchen/mall-api/pkg/code"
	"github.com/yongcycchen/mall-api/vars"
)

func IndexApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, "Welcome to "+vars.App.Name)
	return
}

func PingApi(c *gin.Context) {
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, time.Now())
}
