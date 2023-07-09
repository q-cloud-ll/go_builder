package server

import (
	"fmt"
	"project/setting"
	"project/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"

	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowServer(router *gin.Engine) {
	if err := utils.InitTrans("zh"); err != nil {
		panic(fmt.Errorf("failed to initialize translator: %v", err))
	}
	address := fmt.Sprintf(":%d", setting.Conf.Port)
	s := initServer(address, router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	zap.L().Info("server run success on ", zap.String("address", address))
	csdn := "https://blog.csdn.net/weixin_51991615"
	fmt.Printf(`
	欢迎使用 go_builder
	当前版本:v0.0.1
	简介：主要为了快速搭建小型项目的脚手架
	Up主博客地址：%s
`, csdn)
	zap.L().Error(s.ListenAndServe().Error())
}
