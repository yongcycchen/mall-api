package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/model/args"
	"github.com/yongcycchen/mall-api/pkg/app"
	"github.com/yongcycchen/mall-api/pkg/code"
	"github.com/yongcycchen/mall-api/service"
)

func GetVerifyCodeApi(c *gin.Context) {
	var form args.GenVerifyCodeArgs
	var err error
	err = app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	retCode, verifyCode := service.GenVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, verifyCode)
}
