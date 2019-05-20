package logging

import (
	"context"

	"cloud.google.com/go/logging"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

type StackDriverLogger struct {

}

func NewStackDriverLogger() *StackDriverLogger {
	return &StackDriverLogger{}
}

func (s *StackDriverLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	loggingClient, _ := logging.NewClient(ctx, settings.ProjectName)
	defer loggingClient.Close()
	logger := loggingClient.Logger("Logger").StandardLogger(logging.Info)
	logger.Printf(format, args...)
}

func (s *StackDriverLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	loggingClient, _ := logging.NewClient(ctx, settings.ProjectName)
	defer loggingClient.Close()
	logger := loggingClient.Logger("Logger").StandardLogger(logging.Error)
	logger.Printf(format, args...)
}