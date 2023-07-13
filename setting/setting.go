package setting

import (
	"flag"
	"fmt"
	"os"
	"project/consts"

	"github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name             string `mapstructure:"name"`
	Mode             string `mapstructure:"mode"`
	Version          string `mapstructure:"version"`
	DbType           string `mapstructure:"db-type"`
	Port             int    `mapstructure:"port"`
	*CORS            `mapstructure:"cors"`
	*JWT             `mapstructure:"jwt"`
	*SnowflakeConfig `mapstructure:"snowflake"`
	*LogConfig       `mapstructure:"log"`
	*EsConfig        `mapstructure:"es"`
	*JaegerConfig    `mapstructure:"jaeger"`
	*RabbitMqConfig  `mapstructure:"rabbitmq"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
}

type JaegerConfig struct {
	Addr     string  `mapstructure:"addr"`
	Type     string  `mapstructure:"type"`
	Param    float64 `mapstructure:"param"`
	LogSpans bool    `mapstructure:"log-spans"`
}

type RabbitMqConfig struct {
	RabbitMQ         string `mapstructure:"rabbitmq"`
	RabbitMQUser     string `mapstructure:"rabbitmq_user"`
	RabbitMQPassWord string `mapstructure:"rabbitmq_password"`
	RabbitMQHost     string `mapstructure:"rabbitmq_host"`
	RabbitMQPort     string `mapstructure:"rabbitmq_port"`
}

type EsConfig struct {
	EsHost  string `mapstructure:"es_host"`
	EsPort  string `mapstructure:"es_port"`
	EsIndex string `mapstructure:"es_index"`
}

type JWT struct {
	SigningKey    string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`          // jwt签名
	AccessExpire  string `mapstructure:"access-expire" json:"access-expire" yaml:"access-expire"`    // 缓冲时间
	RefreshExpire string `mapstructure:"refresh-expire" json:"refresh-expire" yaml:"refresh-expire"` // 过期时间
	Issuer        string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                         // 签发者
}

type SnowflakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type CORS struct {
	Mode      string          `mapstructure:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allow-origin"`
	AllowMethods     string `mapstructure:"allow-methods"`
	AllowHeaders     string `mapstructure:"allow-headers"`
	ExposeHeaders    string `mapstructure:"expose-headers"`
	AllowCredentials bool   `mapstructure:"allow-credentials"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Config       string `mapstructure:"config"`
	Port         string `mapstructure:"port"`
	Prefix       string `mapstructure:"prefix"` //全局表前缀，单独定义TableName则不生效
	Engine       string `mapstructure:"engine" default:"InnoDB"`
	LogMode      string `mapstructure:"log_mode"`
	LogZap       bool   `mapstructure:"log_zap"`
	Singular     bool   `mapstructure:"singular"` //是否开启全局禁用复数，true表示开启
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func (m *MySQLConfig) Dsn() string {
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
	cfg, _ := mysql.ParseDSN(dsn)

	return cfg.FormatDSN()
}

type RedisConfig struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	PoolSize     int    `mapstructure:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init(path ...string) (err error) {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(consts.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = consts.ConfigDefaultFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, consts.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = consts.ConfigReleaseFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, consts.ConfigReleaseFile)
				case gin.TestMode:
					config = consts.ConfigTestFile
					fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, consts.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空 将值赋值于config
				config = configEnv
				fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", consts.ConfigEnv, config)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	viper.SetConfigFile(config)

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("天啦噜，配置文件被修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		}
	})

	return
}
