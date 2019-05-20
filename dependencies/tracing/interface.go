package tracing

import (
	"context"
)

type Tracer interface{
	StartSpan(ctx context.Context, spanName string) context.Context
	EndSpan(ctx context.Context)
}