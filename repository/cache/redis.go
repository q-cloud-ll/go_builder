package cache

import (
	"context"
	"os"
	"project/setting"
	"sync"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

var (
	once         sync.Once
	RedisClient  *redis.Client
	RedisContext = context.Background()
)

func InitRedis() {
	rConfig := setting.Conf.RedisConfig
	once.Do(func() {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     rConfig.Addr,
			Password: rConfig.Password,
			DB:       rConfig.DB,
		})
		pong, err := RedisClient.Ping(RedisContext).Result()
		if err != nil {
			zap.L().Error("redis connect ping failed, err:", zap.Error(err))
			os.Exit(0)
			return
		} else {
			zap.L().Info("redis connect ping response:", zap.String("pong", pong))
		}
	})

}
