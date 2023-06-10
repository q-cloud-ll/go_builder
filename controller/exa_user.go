package controller

import (
	"project/model"
	"project/service"
	"project/utils/app"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	var u model.UserParamReq
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("SignUpHandler with param invalid,err:", zap.Error(err))
		app.ResponseError(c, app.CodeInvalidParam)
		return
	}

	if err := service.SignUpService(&u); err != nil {
		zap.L().Error("service.SignUpService(&u) failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err.Error())
		return
	}

	app.ResponseSuccess(c, nil)
}
