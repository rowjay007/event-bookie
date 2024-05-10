package middleware

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "time"
)

// LoggerMiddleware is a middleware to log requests
func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Start timer
        start := time.Now()

        // Process request
        c.Next()

        // Log request details
        latency := time.Since(start)
        clientIP := c.ClientIP()
        method := c.Request.Method
        statusCode := c.Writer.Status()
        path := c.Request.URL.Path

        fmt.Printf("[GIN] %v | %3d | %12v | %s | %-7s %s\n",
            start.Format("2006/01/02 - 15:04:05"),
            statusCode,
            latency,
            clientIP,
            method,
            path,
        )
    }
}
