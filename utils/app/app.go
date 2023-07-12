package app

import (
	"fmt"
	"net/http"
	"project/consts"
	"regexp"

	"github.com/gin-gonic/gin"
)

// ResponseData 返回数据结构体
type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

type TrackedErrorResponse struct {
	ResponseData
	TrackId string `json:"track_id"`
}

// ResponseError 返回错误响应
func ResponseError(c *gin.Context, code ResCode) {
	trackId, _ := getTrackIdFromCtx(c)
	r := &TrackedErrorResponse{
		ResponseData: ResponseData{
			Code: code,
			Msg:  code.Msg(),
			Data: nil,
		},
		TrackId: trackId,
	}
	c.JSON(http.StatusOK, r)
}

// ResponseErrorWithCodeMsg 返回错误响应和信息
func ResponseErrorWithCodeMsg(c *gin.Context, code ResCode, msg interface{}) {
	trackId, _ := getTrackIdFromCtx(c)
	r := &TrackedErrorResponse{
		ResponseData: ResponseData{
			Code: code,
			Msg:  msg,
			Data: nil,
		},
		TrackId: trackId,
	}
	c.JSON(http.StatusOK, r)
}

// ResponseErrorWithMsg 返回错误信息
func ResponseErrorWithMsg(c *gin.Context, msg interface{}) {
	trackId, _ := getTrackIdFromCtx(c)
	r := &TrackedErrorResponse{
		ResponseData: ResponseData{
			Msg:  msg,
			Data: nil,
		},
		TrackId: trackId,
	}
	c.JSON(http.StatusOK, r)
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

func getTrackIdFromCtx(ctx *gin.Context) (trackId string, err error) {
	spanCtxInterface, _ := ctx.Get(consts.SpanCTX)
	str := fmt.Sprintf("%v", spanCtxInterface)
	re := regexp.MustCompile(`([0-9a-fA-F]{16})`)

	match := re.FindStringSubmatch(str)
	if len(match) > 0 {
		return match[1], nil
	}
	return "", consts.GetTrackIdErr
}
