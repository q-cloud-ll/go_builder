package redis

import (
	"context"
	"os"
	"project/initialize/initdb"
	"project/setting"
	"sync"

	"github.com/go-redis/redis/v8"

	"go.uber.org/zap"
)

var (
	once sync.Once
	rdb  *redis.Client
)

func Init(redisCfg *setting.RedisConfig) (err error) {
	once.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			Addr:     redisCfg.Addr,
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		})
		pong, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			zap.L().Error("redis connect ping failed, err:", zap.Error(err))
			os.Exit(0)
			return
		} else {
			zap.L().Info("redis connect ping response:", zap.String("pong", pong))

			initdb.INITPTR[rdb] = struct{}{}
		}
	})

	return
}
