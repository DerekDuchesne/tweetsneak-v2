package tracing

import (
	"context"
)

type MockTracer struct {

}

func NewMockTracer() *MockTracer {
	return &MockTracer{}
}

func (m *MockTracer) StartSpan(ctx context.Context) context.Context {
	return nil
}

func (m *MockTracer) EndSpan(ctx context.Context) {
	return
}