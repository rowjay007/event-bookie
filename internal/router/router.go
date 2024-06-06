package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/config"
	"github.com/rowjay007/event-bookie/internal/handlers"
	"github.com/rowjay007/event-bookie/internal/middleware"
	"github.com/rowjay007/event-bookie/internal/repository"
	"github.com/rowjay007/event-bookie/internal/service"
	"github.com/rowjay007/event-bookie/internal/service/payment"
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

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)

	bookingRepo := repository.NewBookingRepository(db)
	bookingService := service.NewBookingService(bookingRepo)
	bookingHandler := handlers.NewBookingHandler(bookingService)

	paymentRepo := repository.NewPaymentRepository(db)
	conf := config.NewConfig()
	paystackClient := payment.NewPaystackClient(conf) 
	paymentSvc := payment.NewPaymentService(paymentRepo, paystackClient)
	paymentHandler := handlers.NewPaymentHandler(paymentSvc) 


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

		eventGroup := apiV1.Group("/events")
		eventGroup.Use(middleware.AuthMiddleware())
		{
			eventGroup.POST("/", eventHandler.CreateEvent)
			eventGroup.GET("/", eventHandler.GetAllEvents)
			eventGroup.GET("/:id", eventHandler.GetEventByID)
			eventGroup.PUT("/:id", eventHandler.UpdateEvent)
			eventGroup.DELETE("/:id", eventHandler.DeleteEvent)
		}

		bookingGroup := apiV1.Group("/bookings")
		bookingGroup.Use(middleware.AuthMiddleware())
		{
			bookingGroup.GET("/", bookingHandler.GetAllBookings)
			bookingGroup.POST("/", bookingHandler.CreateBooking)
			bookingGroup.GET("/:id", bookingHandler.GetBookingByID)
			bookingGroup.PUT("/:id", bookingHandler.UpdateBooking)
			bookingGroup.DELETE("/:id", bookingHandler.DeleteBooking)
		}

		paymentGroup := apiV1.Group("/payments")
		paymentGroup.Use(middleware.AuthMiddleware())
		{
			paymentGroup.POST("", paymentHandler.CreateAdminPayment)
			paymentGroup.GET("/:id", paymentHandler.GetPaymentByID)
			paymentGroup.PUT("/:id", paymentHandler.UpdatePayment)
			paymentGroup.DELETE("/:id", paymentHandler.DeletePayment)
			paymentGroup.GET("", paymentHandler.GetAllPayments)
		}

		paystackGroup := apiV1.Group("/paystack")
		paystackGroup.Use(middleware.AuthMiddleware())
		{
			paystackGroup.POST("/initialize-payment", paymentHandler.InitializePayment)
			paystackGroup.GET("/verify-payment/:reference", paymentHandler.VerifyPayment)
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
