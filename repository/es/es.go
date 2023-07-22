package es

import (
	"context"
	"fmt"
	"log"
	"project/setting"
	"time"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"

	elastic "github.com/elastic/go-elasticsearch"
	elastic7 "github.com/olivere/elastic/v7"
)

var (
	EsClient *elastic.Client
)

func InitEs() {
	esConn := fmt.Sprintf("http://%s:%s", setting.Conf.EsConfig.EsHost, setting.Conf.EsConfig.EsPort)
	cfg := elastic.Config{
		Addresses: []string{esConn},
	}

	client, err := elastic.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}

	EsClient = client
}

//func EsHookLog() *elogrus.ElasticHook {
//	eConfig := setting.Conf.EsConfig
//	hook, err := elogrus.NewElasticHook(EsClient, eConfig.EsHost, zap.InfoLevel, eConfig.EsIndex)
//	if err != nil {
//		log.Panic(err)
//	}
//	return hook
//}

func EsHookLog() zapcore.Core {
	eConfig := setting.Conf.EsConfig
	client, err := elastic7.NewClient(elastic7.SetURL(eConfig.EsHost))
	if err != nil {
		log.Panic(err)
	}
	esEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	esWriteSyncer := NewElasticsearchWriteSyncer(client, eConfig.EsIndex)

	return zapcore.NewCore(esEncoder, esWriteSyncer, zap.InfoLevel)
}

func NewElasticsearchWriteSyncer(client *elastic7.Client, index string) zapcore.WriteSyncer {
	return &ElasticsearchWriteSyncer{
		client: client,
		index:  index,
	}
}

type ElasticsearchWriteSyncer struct {
	client *elastic7.Client
	index  string
}

func (w ElasticsearchWriteSyncer) Write(p []byte) (n int, err error) {
	//TODO implement me
	ctx := context.Background()
	indexName := w.index + "-" + time.Now().Format("2006.01.02")
	_, err = w.client.Index().
		Index(indexName).
		BodyString(string(p)).
		Do(ctx)
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

func (e ElasticsearchWriteSyncer) Sync() error {
	//TODO implement me
	panic("implement me")
}
