package gcore

import (
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

const (
	serviceName = "HelloClient"
	servName    = "HelloServer"
)
//NewJaegerTracer 初始化调用跟踪
func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := jaegerCfg.Configuration{
		Sampler: &jaegerCfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerCfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  "jaegertracing:6831",
		},
	}
	tracer, closer, err = cfg.New(
		serviceName,
		jaegerCfg.Logger(jaeger.StdLogger),
	)
	//defer closer.Close()

	if err != nil {
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}