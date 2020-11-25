/**
 *@Description
 *@ClassName jaeger
 *@Date 2020/11/24 3:22 下午
 *@Author ckhero
 */

package plugin

import (
	"base-demo/pkg/config"
	"base-demo/pkg/log/logger"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

type Jaeger struct {
	closer  io.Closer
}

func NewPluginJaeger() *Jaeger {
	return &Jaeger{
	}
}

var jaegerTracer opentracing.Tracer

func (e *Jaeger) InitPlugin() error {
	cfg := &jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%s:%d", config.AppConfig.Jaeger.Host, config.AppConfig.Jaeger.Port),
		},
		ServiceName: config.AppConfig.Jaeger.Name,
	}
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	logger.GetLogger(context.TODO()).WithField("address", fmt.Sprintf("%s:%d", config.AppConfig.Jaeger.Host, config.AppConfig.Jaeger.Port)).Info("jaeger start succ")
	if err != nil {
		return err
	}
	jaegerTracer = tracer
	e.closer = closer
	opentracing.SetGlobalTracer(tracer)
	return nil
}

func (e *Jaeger) Release() {
	_ = e.closer.Close()
	fmt.Println("jaeger release")
}

func GetTracer() opentracing.Tracer {
	return jaegerTracer
}