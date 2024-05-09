package middleware

import (
    "net/http"
)

// Example validation middleware
func ExampleValidationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Example validation logic
        // You can implement your own validation logic here
        // For example, you can validate request parameters or headers
        // If validation fails, you can return an error response
        // If validation succeeds, you can call next.ServeHTTP(w, r) to proceed with the next middleware or handler
        next.ServeHTTP(w, r)
    })
}
