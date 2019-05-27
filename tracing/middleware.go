package tracing

import (
	"net/http"

	"go.opencensus.io/trace"
)

// StackDriverTracingMiddleware is middleware that starts a span and ends a StackDriver span around a request being handled.
func StackDriverTracingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := trace.StartSpan(r.Context(), "Search request")
		req := r.WithContext(ctx)
		defer span.End()
		h.ServeHTTP(w, req)
	})
}
