package logging

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/logging"
	"go.opencensus.io/trace"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

var (
	client        *logging.Client
	requestLogger *logging.Logger
	appLogger     *logging.Logger
)

func init() {
	// Create the context.
	ctx := context.Background()

	// Create the logging client.
	c, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", settings.ProjectName))
	if err != nil {
		panic(err)
	}
	client = c

	// Create the loggers.
	requestLogger = client.Logger("request_log")
	appLogger = client.Logger("request_log_entries")
}

// Debugf will log a debug message to StackDriver logging.
func Debugf(ctx context.Context, format string, args ...interface{}) {
	logf(ctx, logging.Debug, format, args...)
}

// Infof will log an info message to StackDriver logging.
func Infof(ctx context.Context, format string, args ...interface{}) {
	logf(ctx, logging.Info, format, args...)
}

// Warningf will log a warning message to StackDriver logging.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	logf(ctx, logging.Warning, format, args...)
}

// Errorf will log an error message to StackDriver logging.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	logf(ctx, logging.Error, format, args...)
}

// logf logs the given message to StackDriver logging.
func logf(ctx context.Context, severity logging.Severity, format string, args ...interface{}) {
	// Increase the severity of the log for the entire request if needed.
	if reqEntry := fromContext(ctx); reqEntry != nil && severity > reqEntry.Severity {
		reqEntry.Severity = severity
	}

	// Create the log entry.
	entry := logging.Entry{
		Severity:  severity,
		Payload:   fmt.Sprintf(format, args...),
		Timestamp: time.Now(),
	}

	// Set the trace ID to the request's trace.
	if span := trace.FromContext(ctx); span != nil {
		entry.Trace = fmt.Sprintf("projects/%s/traces/%s", settings.ProjectName, span.SpanContext().TraceID)
	}

	// Just do a regular log if we're not on production.
	if !settings.UseProduction {
		log.Printf(format, args...)
		return
	}

	// Write the log to StackDriver.
	appLogger.Log(entry)
}
