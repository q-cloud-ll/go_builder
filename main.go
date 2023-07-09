package main

import (
	"fmt"
	"project/dao/redis"
	"project/global"
	"project/initialize"
	"project/initialize/initdb"
	"project/logger"
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
	// 初始化配置
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	initdb.INITPTR = make(map[interface{}]struct{})
	// 初始化日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	// 初始化redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}

	// 初始化gorm
	if err := initialize.Gorm(); err != nil {
		fmt.Printf("init gorm-mysql failed, err:%v\n", err)
		return
	}

	// 初始化雪花算法
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	if initdb.INITPTR != nil {
		initdb.InitDbPtr(initdb.INITPTR)
	}

	if global.GB_MDB != nil {
		initialize.RegisterTables()
		db, _ := global.GB_MDB.DB()

		defer db.Close()
	}

	// 初始化注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	server.RunWindowServer(r)

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}

}
