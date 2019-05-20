package logging

import (
	"context"
	"fmt"
	"log"
)

type MockLogger struct {

}

func NewMockLogger() *MockLogger {
	return &MockLogger{}
}

func (m *MockLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	log.Printf(fmt.Sprintf("INFO: %s", format), args...)
}

func (m *MockLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	log.Printf(fmt.Sprintf("ERROR: %s", format), args...)
}