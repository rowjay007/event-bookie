package middleware

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/rowjay007/event-bookie/internal/utils"
)

// AuthMiddleware is a middleware function to authenticate requests using JWT
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the Authorization header value
        tokenString := c.GetHeader("Authorization")

        // Check if the token is missing
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
            c.Abort()
            return
        }

        // Extract the token from the Authorization header
        parts := strings.Split(tokenString, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }
        token := parts[1]

        // Verify the token
        claims, err := utils.VerifyToken(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Attach the claims to the context for later use
        c.Set("claims", claims)

        // Call the next handler
        c.Next()
    }
}
