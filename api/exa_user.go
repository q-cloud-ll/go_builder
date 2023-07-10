package api

import (
	"project/model"
	systemReq "project/model/request"
	"project/service"
	"project/setting"
	"project/utils"
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

func SignInHandler(c *gin.Context) {
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignInHandler with invalid param", zap.Error(err))
		app.ResponseErrorWithMsg(c, app.CodeInvalidParam)
		return
	}
	if err := service.SignInService(&u); err != nil {
		zap.L().Error("用户名或密码错误,err:", zap.Error(err))
		app.ResponseError(c, app.CodeUserNameOrPasswordFail)
		return
	}
	TokenNext(c, &u)

	return
}

func TokenNext(c *gin.Context, user *model.User) {
	j := utils.JWT{SigningKey: []byte(setting.Conf.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UID:      user.UserId,
		Username: user.UserName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		zap.L().Error("获取token失败", zap.Error(err))
		app.ResponseErrorWithMsg(c, "获取token失败")
		return
	}
	app.ResponseSuccess(c, model.UserSignIn{
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	})
}
