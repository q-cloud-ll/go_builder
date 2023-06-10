package setting

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"

	"time"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}

func RunWindowServer(router *gin.Engine) {
	address := fmt.Sprintf(":%d", Conf.Port)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	zap.L().Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 go_builder
	当前版本:v0.0.1
	简介：主要为了快速搭建小型项目的脚手架
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	zap.L().Error(s.ListenAndServe().Error())
}