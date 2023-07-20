package v1

import (
	"io/ioutil"
	"project/consts"
	"project/utils/app"
	"project/utils/upload"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UploadFileHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")

	if err != nil {
		zap.L().Error("formfile failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err)
		return
	}

	defer file.Close()
	if header.Size > consts.UploadMaxBytes {
		app.ResponseErrorWithMsg(c, "图片不能超过"+strconv.Itoa(consts.UploadMaxM)+"M")
		return
	}
	contentType := header.Header.Get("Content-Type")

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		zap.L().Error("ReadAll failed", zap.Error(err))
		app.ResponseErrorWithMsg(c, err)
		return
	}
	zap.L().Info("上传文件：", zap.String("filename: ", header.Filename), zap.Int64("size:", header.Size))
	if err != nil {
		app.ResponseErrorWithMsg(c, err)
		return
	}
	url, err := upload.PutImage(fileBytes, contentType)
	if err != nil {
		app.ResponseErrorWithMsg(c, err)
		return
	}

	app.ResponseSuccess(c, gin.H{"urls": url})
}
