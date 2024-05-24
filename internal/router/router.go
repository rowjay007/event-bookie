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

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize repository and service layers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Public routes
	r.GET("/", handlers.WelcomeHandler)
	r.POST("/api/v1/signup", userHandler.CreateUser)

	// API v1 routes
	apiV1 := r.Group("/api/v1")
	{
		// User routes
		userGroup := apiV1.Group("/users")
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.GET("/:id", userHandler.GetUserByID)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
			userGroup.GET("/", userHandler.GetAllUsers)
		}

		// Authentication routes
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/login", userHandler.Login)
			authGroup.POST("/forgot-password", userHandler.ForgotPassword)
			authGroup.POST("/reset-password", userHandler.ResetPassword)
			authGroup.POST("/logout", userHandler.Logout)
		}

		// Secured routes
		securedGroup := apiV1.Group("/secured")
		securedGroup.Use(middleware.AuthMiddleware())
		{
			securedGroup.GET("/profile", userHandler.GetUserProfile)
		}
	}

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
