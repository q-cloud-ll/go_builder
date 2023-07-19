## 一、go_builder脚手架介绍

### 1、项目介绍

本脚手架是基于gin框架实现

- 内置zap日志库
- viper读取配置文件
- jwt & cors & track的中间件
- 缓存使用redis，orm框架使用gorm
- 七牛云oss文件上传
- 使用依赖倒置原则抽象dao层
- dockerfile制作镜像，docker-compose编排项目
- 使用shell脚本自动生成dao、cache、service、svc模板代码

准备实现：

- jaeger的链路追踪收集到jaeger-ui
- elk日志收集系统

### 2、项目目的

在使用Gin框架的时候，因为比较轻量级，使用gin.new()，并且run起来就运行起来了，但是我们项目中往往需要配置许多依赖库，我这里就使用cld的架构搭建里一套关于gin框架的脚手架。

### 3、项目介绍

脚手架架构分为CLD分层，controller为api层、service为逻辑层，dao层为数据库层，上手简单，目录结构清晰，一些常用小工具后续会慢慢加上～～～

## 二、go_builder 脚手架目录结构

```shell
├── conf
├── api
├── cmd
├── deploy
├── logger
├── middlewares
├── repository
    ├── cache
    ├── db
      ├── dao
      ├── model
    ├── es
    ├── rabbitmq
    ├── track
├── router
├── service
	  ├── svc
├── setting
└── utils
    ├── app
    └── snowflake
    ├── jwt
    ├── upload
```



| 文件夹        | 说明           | 描述                                                |
| :------------ | -------------- | --------------------------------------------------- |
| `conf`        | 配置包         | 放置配置文件，例：config.yaml                       |
| `api`         | api层          | 程序入口层                                          |
| `cmd`         | 服务程序包     | 包含程序运行main函数以及运行脚本                    |
| `deploy`      | 外来配置工具包 | 配置nginx.conf、sql建表、script脚本、组件yaml配置等 |
| `logger`      | 日志包         | 初始化日志文件                                      |
| `middlewares` | 中间件         | 自定义关于gin的中间件，例如jwt、cors等              |
| `repository`  | 组件仓库       | 包含缓存、db、es、mq等                              |
| `--cache`     | 缓存组件       | 使用redis缓存层                                     |
| `--db`        | sql组件        | 存放mysql的model & dao                              |
| `router`      | 路由层         | 用于放入全局路由                                    |
| `service`     | 逻辑层         | 用于放入业务逻辑                                    |
| `--svc`       | 逻辑上下文层   | 用于初始化service层的context，拿到dao层接口数据     |
| `setting`     | 配置项         | yaml配置映射为结构体                                |
| `utils`       | 工具包         | 自定义工具使用                                      |
| `--app`       | 全局响应       | 返回json数据的封装，success & failed                |
| `--snowflake` | 雪花算法工具包 | 生成int64的id                                       |

### 三、关于gencode自动生成代码

gencode脚本在deploy/gencode/gencode.sh中，首先进入gencode文件目录，如果我要生成关于user的代码，命令：**./gencode.sh user**，会生成对应的svc、service、dao、cache模板代码。后续考虑单独抽出来，放入gopath/bin目录下

### 四、docker的使用

#### 1、使用dockerfile制作镜像

`docker build -t go_builder .` 构建镜像

#### 2、使用docker-compose一键编排项目

`docekr-compose up` 一键启动项目

`docker-compose stop` 一键关闭项目



## 欢迎大家提issue！！！