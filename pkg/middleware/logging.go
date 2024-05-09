package middleware

import (
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware logs information about incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        defer func() {
            log.Printf("[%s] %s %s %v", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
        }()
        next.ServeHTTP(w, r)
    })
}
