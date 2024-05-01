package middleware

import "net/http"

// ValidationMiddleware performs request validation
func ValidationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add validation logic here

        // Serve the next handler
        next.ServeHTTP(w, r)
    })
}
