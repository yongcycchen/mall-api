package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yongcycchen/mall-api/pkg/app"
	"github.com/yongcycchen/mall-api/pkg/code"
)

func CheckUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenEmpty, code.GetMsg(code.ErrorTokenEmpty))
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenInvalid, code.GetMsg(code.ErrorTokenInvalid))
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			app.JsonResponse(c, http.StatusOK, code.ErrorTokenExpire, code.GetMsg(code.ErrorTokenExpire))
			c.Abort()
			return
		}
		if claims == nil || claims.Uid == 0 {
			app.JsonResponse(c, http.StatusOK, code.ErrorUserNotExist, code.GetMsg(code.ErrorUserNotExist))
			c.Abort()
			return
		}
	}
}
