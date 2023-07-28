package kafka

import (
	"context"
	"strings"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// Consumer 消费者函数
func Consumer(ctx context.Context, key, topic string, fn func(message *sarama.ConsumerMessage) error) (err error) {
	kafka, err := GetClient(key)
	if err != nil {
		return
	}
	partitions, err := kafka.Consumer.Partitions(topic)
	if err != nil {
		return
	}
	for _, partition := range partitions {
		// 针对每个分区创建一个对应的分区消费者
		offset, errx := kafka.Client.GetOffset(topic, partition, sarama.OffsetNewest)
		if errx != nil {
			zap.L().Info("获取Offset失败:", zap.Error(errx))
			continue
		}
		pc, errx := kafka.Consumer.ConsumePartition(topic, partition, offset-1)
		if errx != nil {
			zap.L().Info("获取Offset失败:", zap.Error(errx))
			return nil
		}
		// 从每个分区都消费消息
		go func(consume sarama.PartitionConsumer) {
			defer func() {
				if err := recover(); err != nil {
					zap.L().Error("消费kafka信息发生panic,err:%s", zap.Any("err:", err))
				}
			}()

			defer func() {
				err := pc.Close()
				if err != nil {
					zap.L().Error("消费kafka信息发生panic,err:%s", zap.Any("err:", err))
				}
			}()

			for {
				select {
				case msg := <-pc.Messages():
					err := MiddlewareConsumerHandler(fn)(msg)
					if err != nil {
						return
					}
				case <-ctx.Done():
					return
				}
			}

		}(pc)
	}
	return nil
}

// ConsumerGroup 消费者组消费消息
func ConsumerGroup(ctx context.Context, key, groupId, topics string, msgHandler ConsumerGroupHandler) (err error) {
	kafka, err := GetClient(key)
	if err != nil {
		return
	}

	if isConsumerDisabled(kafka) {
		return
	}

	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupId, kafka.Client)
	if err != nil {
		return
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("消费kafka发生panic", zap.Any("panic", err))
			}
		}()

		defer func() {
			err := consumerGroup.Close()
			if err != nil {
				zap.L().Error("close err", zap.Any("panic", err))
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := consumerGroup.Consume(ctx, strings.Split(topics, ","), ConsumerGroupHandler(func(msg *sarama.ConsumerMessage) error {
					return MiddlewareConsumerHandler(msgHandler)(msg)
				}))
				if err != nil {
					zap.L().Error("消费kafka失败 err", zap.Any("panic", err))

				}
			}
		}
	}()

	return
}

func isConsumerDisabled(in *Kafka) bool {
	if in.DisableConsumer {
		zap.L().Info("kafka consumer disabled,key:%s", zap.String("key", in.Key))
	}

	return in.DisableConsumer
}

func MiddlewareConsumerHandler(fn func(message *sarama.ConsumerMessage) error) func(message *sarama.ConsumerMessage) error {
	return func(msg *sarama.ConsumerMessage) error {
		return fn(msg)
	}
}

type ConsumerGroupHandler func(message *sarama.ConsumerMessage) error

func (ConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (ConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h ConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := h(msg); err != nil {
			zap.L().Info("消息处理失败",
				zap.String("topic", msg.Topic),
				zap.String("value", string(msg.Value)))
			continue
		}
		sess.MarkMessage(msg, "")
	}

	return nil
}
