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

// ValidationMiddleware is a middleware for request validation
func ValidationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add validation logic here
        // Example: Validate request parameters, headers, etc.
        // If validation fails, return an appropriate response (e.g., 400 Bad Request)
        // If validation passes, proceed with the next handler
        next.ServeHTTP(w, r)
    })
}
