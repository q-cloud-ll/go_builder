package cache

import (
	"context"
	"os"
	"project/setting"
	"sync"

	logging "github.com/sirupsen/logrus"

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
			logging.Error("redis connect ping failed, err:", zap.Error(err))
			os.Exit(0)
			return
		} else {
			logging.Info("redis connect ping response:", zap.String("pong", pong))
		}
	})

}
