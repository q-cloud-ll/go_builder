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
	once sync.Once
	Rdb  *redis.Client
)

func InitRedis(redisCfg *setting.RedisConfig) {
	once.Do(func() {
		Rdb = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
		pong, err := Rdb.Ping(context.Background()).Result()
		if err != nil {
			logging.Error("redis connect ping failed, err:", zap.Error(err))
			os.Exit(0)
			return
		} else {
			logging.Info("redis connect ping response:", zap.String("pong", pong))
		}
	})

}
