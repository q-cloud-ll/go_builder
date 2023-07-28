package kafka

import (
	"fmt"
	"project/setting"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/IBM/sarama"
)

type Kafka struct {
	Key             string
	DisableConsumer bool
	Debug           bool
	Producer        sarama.SyncProducer
	Consumer        sarama.Consumer
	Client          sarama.Client
}

var kafkaClient = new(sync.Map)

func InitKafka() {
	for k, v := range setting.Conf.KafKa {
		key := k
		val := v
		scfg := buildConfig(val)
		kafka, err := newKafkaClient(key, val, scfg)
		if err != nil {
			zap.L().Error("newKafkaClient(key, val, scfg) failed:", zap.Error(err))
			return
		}
		kafkaClient.Store(key, kafka)
	}
}

func GetClient(key string) (*Kafka, error) {
	val, ok := kafkaClient.Load(key)
	if !ok {
		return nil, fmt.Errorf("获取kafka client失败，key：%s", key)
	}

	return val.(*Kafka), nil
}

func buildConfig(v *setting.KafkaConfig) *sarama.Config {
	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.RequiredAcks(v.RequiredAck)
	cfg.Producer.Return.Successes = true

	if v.Partition == 1 {
		cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	}

	if v.Partition == 2 {
		cfg.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	}

	if v.ReadTimeout != 0 {
		cfg.Net.ReadTimeout = time.Duration(v.ReadTimeout) * time.Second
	}

	if v.WriteTimeout != 0 {
		cfg.Net.WriteTimeout = time.Duration(v.WriteTimeout) * time.Second
	}

	if v.MaxOpenRequests != 0 {
		cfg.Net.MaxOpenRequests = v.MaxOpenRequests
	}

	return cfg
}

func newKafkaClient(key string, cfg *setting.KafkaConfig, scfg *sarama.Config) (*Kafka, error) {
	client, err := sarama.NewClient(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	syncProducer, err := sarama.NewSyncProducer(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	consumer, err := sarama.NewConsumer(strings.Split(cfg.Address, ","), scfg)
	if err != nil {
		return nil, err
	}

	return &Kafka{
		Key:             key,
		DisableConsumer: cfg.DisableConsumer,
		Debug:           cfg.Debug,
		Producer:        syncProducer,
		Consumer:        consumer,
		Client:          client,
	}, nil

}
