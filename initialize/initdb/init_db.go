package initdb

import (
	"project/global"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-redis/redis/v8"
)

var INITPTR map[interface{}]struct{}

func InitDbPtr(dbList map[interface{}]struct{}) {
	for k, _ := range dbList {
		switch k.(type) {
		case *redis.Client:
			global.GB_RDB = k.(*redis.Client)
		case *gorm.DB:
			global.GB_MDB = k.(*gorm.DB)
		case *zap.Logger:
			global.GB_LOG = k.(*zap.Logger)
		default:
			zap.L().Error("初始化全局指针失败", zap.Any("ptr:", k))
		}
	}
}
