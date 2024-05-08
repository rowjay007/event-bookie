package middleware

import (
	"github.com/gin-gonic/gin"
)

// ValidateInput is a middleware function to validate request input
func ValidateInput() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Example: Validate request body here

        // Call the next handler
        c.Next()
    }
}
