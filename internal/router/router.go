package router

import (
    "github.com/gin-gonic/gin"
    // "github.com/rowjay007/event-bookie/internal/handlers"
    "github.com/rowjay007/event-bookie/internal/middleware"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
   
)

// NewRouter creates a new router instance with defined routes
func NewRouter() *gin.Engine {
    r := gin.Default()

    // Middleware
    r.Use(middleware.LoggerMiddleware())

    // Routes
    // v1 := r.Group("/api/v1")
    // {
    //     // Event routes
    //     eventGroup := v1.Group("/events")
    //     {
    //         eventGroup.GET("/", handlers.GetAllEvents)
    //         eventGroup.POST("/", handlers.CreateEvent)
    //         eventGroup.GET("/:id", handlers.GetEventByID)
    //         // Add more event routes as needed
    //     }

    //     // User routes
    //     userGroup := v1.Group("/users")
    //     {
    //         userGroup.GET("/", handlers.GetAllUsers)
    //         userGroup.POST("/", handlers.CreateUser)
    //         userGroup.GET("/:id", handlers.GetUserByID)
    //         // Add more user routes as needed
    //     }

    //     // Add more route groups for other entities like bookings, venues, payments, organizers, etc.
    // }

    // Swagger routes
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
