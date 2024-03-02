package tracing

import (
	"context"
	"github.com/oigi/Magikarp/config"
	"go.uber.org/zap"

	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
)

func InitTracerProvider(url string, serviceName string) func(ctx context.Context) error {
	ctx := context.Background()
	// 创建一个新的 OTLP gRPC 客户端
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(url),
	)
	// 创建一个新的 OTLP 导出器
	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		config.LOG.Error("failed to init tracer, err: ", zap.Error(err))
		return nil
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exporter),                  //注册exporter
		tracesdk.WithResource(newResource(serviceName)), //设置服务信息
	)
	//设置全局tracer
	otel.SetTracerProvider(tp)
	b3Propagator := b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader))
	propagator := propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}, b3Propagator)
	// 设置全局Propagator
	otel.SetTextMapPropagator(propagator)
	return tp.Shutdown
}

func newResource(serviceName string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
	)
}

func GetTraceID(ctx context.Context) string {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.HasTraceID() {
		traceID := spanCtx.TraceID()
		return traceID.String()
	}

	return ""
}
