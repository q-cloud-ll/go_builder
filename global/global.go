package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GB_MDB *gorm.DB
	GB_RDB *redis.Client
	GB_LOG *zap.Logger
)
