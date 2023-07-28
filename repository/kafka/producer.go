package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// SendMessage 发送消息，默认分区
func SendMessage(key, topic, value string) error {
	return SendMessagePartitionPar(key, topic, value, "")
}

// SendMessagePartitionPar 发送消息，指定分区
func SendMessagePartitionPar(key, topic, value, partitionKey string) error {
	kafka, err := GetClient(key)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(value),
		Timestamp: time.Now(),
	}

	if partitionKey != "" {
		msg.Key = sarama.StringEncoder(partitionKey)
	}
	partition, offset, err := kafka.Producer.SendMessage(msg)
	if err != nil {
		return err
	}
	if kafka.Debug {
		zap.L().Info("发送kafka消息成功",
			zap.Int32("partition", partition),
			zap.Int64("offset", offset))
	}

	return err
}
