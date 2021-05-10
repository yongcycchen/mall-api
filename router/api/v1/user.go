package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/model/args"
	"github.com/yongcycchen/mall-api/pkg/app"
	"github.com/yongcycchen/mall-api/pkg/code"
	"github.com/yongcycchen/mall-api/service"
)

func RegisterUserApi(c *gin.Context) {
	var form args.RegisterUserArgs
	// var err error
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	rsp, retCode := service.CreateUser(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, rsp)
}

func GetVerifyCodeApi(c *gin.Context) {
	var form args.GenVerifyCodeArgs
	// var err error
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	retCode, verifyCode := service.GenVerifyCode(c, &form)
	app.JsonResponse(c, http.StatusOK, retCode, verifyCode)
}

func LoginUserWithVerifyCodeApi(c *gin.Context) {
	var form args.LoginUserWithVerifyCodeArgs
	// var err error
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	token, retCode := service.LoginUserWithVerifyCode(c, &form)
	if retCode != code.SUCCESS {
		app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
		return
	}
	c.Writer.Header().Add("token", token)
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, token)
}

func LoginUserWithPwdApi(c *gin.Context) {
	var form args.LoginUserWithPwdArgs
	// var err error
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.JsonResponse(c, http.StatusOK, code.InvalidParams, err.Error())
		return
	}
	token, retCode := service.LoginUserWithPwd(c, &form)
	if retCode != code.SUCCESS {
		app.JsonResponse(c, http.StatusOK, retCode, code.GetMsg(retCode))
		return
	}
	c.Writer.Header().Add("token", token)
	app.JsonResponse(c, http.StatusOK, code.SUCCESS, token)
}
