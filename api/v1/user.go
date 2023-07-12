package v1

import (
	"project/service"
	"project/types"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRegisterHandler(c *gin.Context) {
	var u types.UserRegisterReq
	if err := c.ShouldBind(&u); err != nil {
		zap.L().Error("UserRegisterHandler param with invalid", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	l := service.GetUserSrv()
	if err := l.UserRegisterSrv(c.Request.Context(), &u); err != nil {
		zap.L().Error("UserRegisterSrv failed,err:", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, nil)
}
