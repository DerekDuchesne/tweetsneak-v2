package logging

import (
	"context"

	"cloud.google.com/go/logging"
)

// loggingKey is the key used to store a request's log entry in a context.
type loggingKey struct{}

// withContext adds the given log entry to the context.
func withContext(ctx context.Context, entry *logging.Entry) context.Context {
	return context.WithValue(ctx, loggingKey{}, entry)
}

// fromContext gets the log entry from the context.
func fromContext(ctx context.Context) *logging.Entry {
	entry, ok := ctx.Value(loggingKey{}).(*logging.Entry)
	if !ok {
		return nil
	}
	return entry
}
