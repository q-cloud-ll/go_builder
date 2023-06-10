package main

import (
	"fmt"
	"go_builder/setting"
)

// @title go_builder
// @version 1.0
// @description 基于Go Web 简易脚手架

// @contact.name camellia
// @contact.url https://github.com/q-cloud-ll

// @host 127.0.0.1:8888
// @BasePath /api/v1

func main() {
	if err := setting.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
}
