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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	organizerRepo := repository.NewOrganizerRepository(db)
	organizerService := service.NewOrganizerService(organizerRepo)
	organizerHandler := handlers.NewOrganizerHandler(organizerService)

	venueRepo := repository.NewVenueRepository(db)
	venueService := service.NewVenueService(venueRepo)
	venueHandler := handlers.NewVenueHandler(venueService)

	r.GET("/", handlers.WelcomeHandler)
	r.POST("/api/v1/signup", userHandler.CreateUser)

	apiV1 := r.Group("/api/v1")
	{
		userGroup := apiV1.Group("/users")
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.GET("/:id", userHandler.GetUserByID)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
			userGroup.GET("/", userHandler.GetAllUsers)
		}

		categoryGroup := apiV1.Group("/categories")
		categoryGroup.Use(middleware.AuthMiddleware())
		{
			categoryGroup.GET("/", categoryHandler.GetAllCategories)
			categoryGroup.POST("/", categoryHandler.CreateCategory)
			categoryGroup.GET("/:id", categoryHandler.GetCategoryByID)
			categoryGroup.PUT("/:id", categoryHandler.UpdateCategory)
			categoryGroup.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		venueGroup := apiV1.Group("/venues")
		venueGroup.Use(middleware.AuthMiddleware())
		{
			venueGroup.POST("/", venueHandler.CreateVenue)
			venueGroup.GET("/", venueHandler.GetAllVenues)
			venueGroup.GET("/:id", venueHandler.GetVenueByID)
			venueGroup.PUT("/:id", venueHandler.UpdateVenue)
			venueGroup.DELETE("/:id", venueHandler.DeleteVenue)
		}


		organizerGroup := apiV1.Group("/organizers")
		organizerGroup.Use(middleware.AuthMiddleware())
		{
			organizerGroup.GET("/", organizerHandler.GetAllOrganizers)
				organizerGroup.POST("/", organizerHandler.CreateOrganizer)
				organizerGroup.GET("/:id", organizerHandler.GetOrganizerByID)
				organizerGroup.PUT("/:id", organizerHandler.UpdateOrganizer)
				organizerGroup.DELETE("/:id", organizerHandler.DeleteOrganizer)
		}

		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/login", userHandler.Login)
			authGroup.POST("/forgot-password", userHandler.ForgotPassword)
			authGroup.POST("/reset-password", userHandler.ResetPassword)
			authGroup.POST("/logout", userHandler.Logout)
		}

		securedGroup := apiV1.Group("/secured")
		securedGroup.Use(middleware.AuthMiddleware())
		{
			securedGroup.GET("/profile", userHandler.GetUserProfile)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
