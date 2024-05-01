package middleware

import (
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware logs incoming HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Serve the next handler
        next.ServeHTTP(w, r)

        // Log request details
        log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
    })
}
