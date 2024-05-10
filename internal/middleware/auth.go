package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// AuthMiddleware is a middleware to authenticate requests
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Check if authentication token is present in the request header
        token := c.GetHeader("Authorization")
        if token == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        // Your authentication logic here (e.g., verify JWT token)

        // Call the next handler
        c.Next()
    }
}
