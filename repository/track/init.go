package track

import (
	"fmt"
	"io"
	"net/http"
	"project/setting"

	"github.com/opentracing/opentracing-go/ext"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func GetDefaultConfig() *config.Configuration {
	jConfig := setting.Conf.JaegerConfig
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jConfig.Type,
			Param: jConfig.Param,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           jConfig.LogSpans,
			LocalAgentHostPort: jConfig.Addr,
		},
	}

	return cfg
}

func InitJaeger() (opentracing.Tracer, io.Closer) {
	cfg := GetDefaultConfig()
	cfg.ServiceName = setting.Conf.Name
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: connot init Jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

func GetParentSpan(spanName string, traceId string, header http.Header) (opentracing.Span, error) {
	carrier := opentracing.HTTPHeadersCarrier{}
	carrier.Set("uber-trace-id", traceId)

	tracer := opentracing.GlobalTracer()
	wireContext, err := tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	parentSpan := opentracing.StartSpan(
		spanName,
		ext.RPCServerOption(wireContext),
	)

	if err != nil {
		return nil, err
	}
	return parentSpan, err
}

func StartSpan(tracer opentracing.Tracer, name string) opentracing.Span {
	// 设置顶级span
	span := tracer.StartSpan(name)
	return span
}
