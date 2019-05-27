package logging

import (
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/logging"
	"github.com/go-chi/chi/middleware"
	"go.opencensus.io/trace"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

// StackDriverLoggingMiddleware is a middleware that will create a log entry for a request.
func StackDriverLoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start the clock for how long the request will take.
		start := time.Now()

		// Create the log entry for the request.
		reqEntry := &logging.Entry{
			HTTPRequest: &logging.HTTPRequest{
				Request:  r,
				RemoteIP: r.RemoteAddr,
			},
		}

		// Set the span and trace if they're available from the context (should be set by the tracing middleware).
		if span := trace.FromContext(r.Context()); span != nil {
			reqEntry.Trace = fmt.Sprintf("projects/%s/traces/%s", settings.ProjectName, span.SpanContext().TraceID)
			reqEntry.SpanID = fmt.Sprintf("%s", span.SpanContext().SpanID)
		}

		// Add the request log entry to the context so that it can be used later in the request.
		req := r.WithContext(withContext(r.Context(), reqEntry))

		// Wrap the response writer so that we can keep track of what it writes.
		resp := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		// After the request is done, write the log.
		defer func() {
			reqEntry.HTTPRequest.Status = resp.Status()
			reqEntry.HTTPRequest.ResponseSize = int64(resp.BytesWritten())
			reqEntry.Timestamp = time.Now()
			reqEntry.HTTPRequest.Latency = reqEntry.Timestamp.Sub(start)
			if settings.UseProduction {
				requestLogger.Log(*reqEntry)
			}
		}()

		// Run the request handler.
		h.ServeHTTP(resp, req)
	})
}
