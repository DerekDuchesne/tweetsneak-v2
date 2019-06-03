package handlers

import (
	"net/http"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

// CorsMiddleware will set headers allowing cross-origin requests so that the client and server can talk to each other from different ports when not on production.
func CorsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !settings.UseProduction {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Header().Set("Access-Control-Allow-Methods", "GET")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		h.ServeHTTP(w, r)
	})
}
