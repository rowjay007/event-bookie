package middleware

import (
    "net/http"
)

// AuthMiddleware is a middleware for authentication
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        
       
        next.ServeHTTP(w, r)
    })
}
