package tracing

import (
	"context"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

type StackDriverTracer struct {

}

func init() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{})
	if err != nil {
		panic(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func NewStackDriverTracer()  *StackDriverTracer {
	return &StackDriverTracer{}
}

func (s *StackDriverTracer) StartSpan(ctx context.Context, spanName string) context.Context {
	ctx, _ = trace.StartSpan(ctx, spanName)
	return ctx
}

func (s *StackDriverTracer) EndSpan(ctx context.Context) {
	span := trace.FromContext(ctx)
	span.End()
}