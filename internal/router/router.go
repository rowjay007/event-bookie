package router

import (
    "github.com/gin-gonic/gin"
    "github.com/rowjay007/event-bookie/internal/handlers"
    "github.com/rowjay007/event-bookie/internal/middleware"
    "github.com/rowjay007/event-bookie/internal/repository"
    "github.com/rowjay007/event-bookie/internal/service"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "gorm.io/gorm"
)

// NewRouter sets up the router with routes and middleware
func NewRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // Apply LoggerMiddleware globally
    router.Use(middleware.LoggerMiddleware())

    // Initialize the repository
    userRepository := repository.NewUserRepository(db)

    // Public routes
    v1 := router.Group("/api/v1")
    {
        v1.GET("/", handlers.WelcomeHandler)

        // User routes (without AuthMiddleware)
        userService := service.NewUserService(userRepository)
        userHandler := handlers.NewUserHandler(userService)
        v1.POST("/users", userHandler.CreateUser)
        v1.GET("/users/:id", userHandler.GetUserByID)
        v1.PUT("/users/:id", userHandler.UpdateUser)
        v1.DELETE("/users/:id", userHandler.DeleteUser)

        // Define your other unprotected routes here
    }

    // Protected routes
    protected := router.Group("/api/v1")
    protected.Use(middleware.AuthMiddleware())
    {
        // Define your protected routes here
    }

    admin := router.Group("/admin")
    admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
    {
        // Define your admin routes here
    }

    // Setup Swagger
    SetupSwagger(router)

    return router
}

// SetupSwagger sets up the Swagger documentation routes
func SetupSwagger(r *gin.Engine) {
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.StaticFile("/swagger.json", "./docs/swagger.json")
}
