package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseData 返回数据结构体
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

// ResponseError 返回错误响应
func ResponseError(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	})
}

// ResponseErrorWithCodeMsg 返回错误响应和信息
func ResponseErrorWithCodeMsg(c *gin.Context, code ResCode, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}

// ResponseErrorWithMsg 返回错误信息
func ResponseErrorWithMsg(c *gin.Context, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Msg:  msg,
		Data: nil,
	})
}

// ResponseSuccess 返回成功响应
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

// ResponseSuccessWithMsg 返回成功，自带文字描述
func ResponseSuccessWithMsg(c *gin.Context, data interface{}, msg interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: CodeSuccess,
		Msg:  msg,
		Data: data,
	})
}
