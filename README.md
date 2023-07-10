## 一、go_builder脚手架介绍

### 1、项目技术使用

**gin+gorm+redis**，后续会引入需要的技术

### 2、项目目的

当我们有一个新的**idea**需要马上付出实践，用于构建小型项目，直接上手写接口即可，主要为了大学生可以快速完成作业，不需要搭建环境，本项目暂时完成不了复杂的业务哦～

### 3、项目介绍

脚手架架构分为CLD分层，controller为api层、service为逻辑层，dao层为数据库层，上手简单，目录结构清晰，一些常用小工具后续会慢慢加上～～～

## 二、go_builder 脚手架目录结构

```shell
├── conf
├── api
├── dao
├── deploy
├── global
├── initialize
    └── initdb
├── logger
├── middlewares
├── model
│   ├── request
├── router
├── service
├── setting
└── utils
    ├── app
    └── snowflake
```



| 文件夹           | 说明      | 描述                           |
|---------------|---------|------------------------------|
| `conf`        | 配置包     | 放置配置文件，例：config.yaml         |
| `controller`  | api层    | 程序入口层                        |
| `dao`         | dao层    | 数据层，操作mysql及redis            |
| `deploy`      | 外来配置工具包 | 配置nginx.conf、sql建表、script脚本等 |
| `global`      | 全局化变量   | 配置全局变量、sql、redis等            |
| `initialize`  | 初始化数据   | 初始化全局所需要的数据                  |
| `logger`      | 日志包     | 初始化日志文件                      |
| `middlewares` | 中间件     | 自定义关于gin的中间件，例如jwt、cors等     |
| `middleware`  | 中间件层    | 用于存放 `gin` 中间件代码             |
| `model`       | 模型层     | 入参出参对应的struct、表对应的struct     |
| `--request`   | 入参结构体   | 系统结构入参、业务入参出参数放入一个文件即可       |
| `router`      | 路由层     | 用于放入全局路由                     |
| `service`     | 逻辑层     | 用于放入业务逻辑                     |
| `setting`     | 配置项     | yaml配置映射为结构体                 |
| `utils`       | 工具包     | 自定义工具使用                      |
| `--app`       | 全局响应    | 返回json数据的封装，success & failed |
| `--snowflake` | 雪花算法工具包 | 生成int64的id                   |

## 三、docker的使用
#### 1、使用dockerfile制作镜像
`docker build -t go_builder .` 构建镜像
#### 2、使用docker-compose一键编排项目
`docekr-compose up` 一键启动项目 

`docker-compose stop` 一键关闭项目



## 欢迎大家提issue！！！