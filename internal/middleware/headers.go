package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// CacheControl returns middleware that sets HTTP Cache-Control headers on GET responses.
// This instructs browsers and CDNs to cache responses for the specified duration.
func CacheControl(maxAge time.Duration) func(http.Handler) http.Handler {
	directive := fmt.Sprintf("public, max-age=%d", int(maxAge.Seconds()))

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodGet {
				w.Header().Set("Cache-Control", directive)
			}
			next.ServeHTTP(w, r)
		})
	}
}

// SecurityHeaders is middleware that sets common security-related HTTP headers
// to help protect against XSS, clickjacking, and MIME sniffing attacks.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}
