package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/handlers"
	"github.com/rowjay007/event-bookie/internal/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files" 
)

// NewRouter initializes and returns a new Gin router
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Load Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Use logger middleware
	r.Use(middleware.LoggerMiddleware())

	// Define API routes with auth middleware
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware()) 
	{
		v1.GET("/", handlers.WelcomeHandler)
		// v1.GET("/events", handlers.GetEvents)
		// v1.POST("/events", handlers.CreateEvent)
	
	}

	// Handle not found routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	return r
}
