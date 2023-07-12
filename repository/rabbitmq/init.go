package rabbitmq

import (
	"project/setting"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

// GlobalRabbitMQ rabbitMQ链接单例
var GlobalRabbitMQ *amqp.Connection

// InitRabbitMQ 在中间件中初始化rabbitMQ链接
func InitRabbitMQ() {
	rConfig := setting.Conf.RabbitMqConfig
	pathRabbitMQ := strings.Join([]string{rConfig.RabbitMQ, "://", rConfig.RabbitMQUser, ":", rConfig.RabbitMQPassWord, "@", rConfig.RabbitMQHost, ":", rConfig.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ)
	if err != nil {
		panic(err)
	}
	GlobalRabbitMQ = conn
}
