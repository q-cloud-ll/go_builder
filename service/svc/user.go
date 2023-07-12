package svc

import (
	"project/repository/cache"
	"project/repository/db/dao"

	"github.com/go-redis/redis/v8"
)

type UserServiceContext struct {
	RedisClient *redis.Client

	UserModel dao.UserModel
}

func NewUserServiceContext() *UserServiceContext {
	return &UserServiceContext{
		RedisClient: cache.RedisClient,
		UserModel:   dao.NewUserModel(),
	}
}
