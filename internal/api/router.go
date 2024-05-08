package api

import (
    "database/sql" 
    "github.com/gin-gonic/gin"
    "github.com/rowjay007/event-bookie/internal/api/middleware"
)

// RegisterRoutes registers API routes
func RegisterRoutes(router *gin.Engine, db *sql.DB) {
    // Apply middleware
    router.Use(middleware.AuthMiddleware())
    router.Use(middleware.ValidateInput())

    // Define API routes here
    router.GET("/api/resource", getResourceHandler)
    router.POST("/api/resource", createResourceHandler)
    router.PUT("/api/resource/:id", updateResourceHandler)
    router.DELETE("/api/resource/:id", deleteResourceHandler)
}

func getResourceHandler(c *gin.Context) {
    // Handle GET request for resource
}

func createResourceHandler(c *gin.Context) {
    // Handle POST request to create resource
}

func updateResourceHandler(c *gin.Context) {
    // Handle PUT request to update resource
}

func deleteResourceHandler(c *gin.Context) {
    // Handle DELETE request to delete resource
}
