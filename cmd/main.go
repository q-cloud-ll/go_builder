package main

import (
	"fmt"
	"project/logger"
	"project/repository/cache"
	"project/repository/db/dao"
	"project/repository/es"
	"project/repository/track"
	"project/router"
	"project/setting"
	"project/setting/server"
	"project/utils/snowflake"
)

// @title go_builder
// @version 1.0
// @description 基于Go Web 简易脚手架

// @contact.name camellia
// @contact.url https://github.com/q-cloud-ll

// @host 127.0.0.1:8888
// @BasePath /api/v1

func main() {
	loadingConfig()
	// 初始化注册路由
	r := router.SetupRouter()
	server.RunWindowServer(r)
	fmt.Println("Starting configuration success...")
	_ = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
}

func loadingConfig() {
	setting.Init()
	logger.Init()
	dao.InitMysql()
	cache.InitRedis()
	es.InitEs()
	//rabbitmq.InitRabbitMQ()
	track.InitJaeger()
	snowflake.InitSnowflake()
	fmt.Println("Loading configuration success...")
	go scriptStarting()
}

func scriptStarting() {
	// start script
}
