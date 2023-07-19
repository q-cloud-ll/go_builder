#!/bin/bash

# 输入命令，such as ./gencode.sh comment
template_name=$1

# 定义目录路径数组
dir_paths=(
  "../../repository/db/dao/"
  "../../repository/cache/"
  "../../service/svc/"
  "../../service/"
)

# 循环遍历目录路径数组
for dir_path in "${dir_paths[@]}"; do
  # 检查目录是否存在，如果不存在则创建目录
  if [ ! -d "$dir_path" ]; then
    mkdir -p "$dir_path"
  fi
done


# 将输入的comment更改为Comment
template_name_upper="$(tr '[:lower:]' '[:upper:]' <<< "${template_name:0:1}")${template_name:1}"

# dao包下模板
template_dao="package dao

import \"gorm.io/gorm\"

var _ ${template_name_upper}Model = (*custom${template_name_upper}Model)(nil)

type (
	  // ${template_name_upper}Model is an interface to be customized, add more methods here,
	  // and implement the added methods in custom${template_name_upper}Model.
	  ${template_name_upper}Model interface {
	  }

	  custom${template_name_upper}Model struct {
		  *gorm.DB
	  }
)

func New${template_name_upper}Model() ${template_name_upper}Model {
	  return &custom${template_name_upper}Model{
		  DB: NewDBClient(),
	  }
}
"

echo "$template_dao" > "${dir_paths[0]}/${template_name}.go"

# cache包下模板
template_cache="package cache

import (
	\"context\"
	\"github.com/go-redis/redis/v8\"
)

var _ ${template_name_upper}Cache = (*custom${template_name_upper}Cache)(nil)

type (
	  // ${template_name_upper}Cache is an interface to be customized, add more methods here,
	  // and implement the added methods in custom${template_name_upper}Cache.
	  ${template_name_upper}Cache interface {
	  }
	  custom${template_name_upper}Cache struct {
		  *redis.Client
	  }
)

func New${template_name_upper}Cache() ${template_name_upper}Cache {
	  return &custom${template_name_upper}Cache{
		  Client: RedisClient,
	  }
}
"
echo "$template_cache" > "${dir_paths[1]}/${template_name}.go"


# svc包下模板
template_svc="package svc

import (
	\"project/repository/cache\"
	\"project/repository/db/dao\"
)

type ${template_name_upper}ServiceContext struct {
    ${template_name_upper}Model dao.${template_name_upper}Model
	  ${template_name_upper}Cache cache.${template_name_upper}Cache
}

func New${template_name_upper}ServiceContext() *${template_name_upper}ServiceContext {
	  return &${template_name_upper}ServiceContext{
		  ${template_name_upper}Model: dao.New${template_name_upper}Model(),
		  ${template_name_upper}Cache: cache.New${template_name_upper}Cache(),
	  }
}
"
echo "$template_svc" > "${dir_paths[2]}/${template_name}.go"

# service包下模板
template_service="package service

import (
    \"context\"
    \"go.uber.org/zap\"
    \"project/logger\"
    \"project/service/svc\"
)


type ${template_name_upper}Srv struct {
    ctx    context.Context
    svcCtx *svc.${template_name_upper}ServiceContext
    log    *zap.Logger
}

func New${template_name_upper}Service(ctx context.Context, svcCtx *svc.${template_name_upper}ServiceContext) *${template_name_upper}Srv {
    return &${template_name_upper}Srv{
        ctx:    ctx,
        svcCtx: svcCtx,
        log:    logger.Lg,
      }
}"

 # 生成模板代码
echo "$template_service" > "${dir_paths[3]}/${template_name}.go"
