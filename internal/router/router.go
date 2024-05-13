package router

import (
    "github.com/gin-gonic/gin"
    // "github.com/rowjay007/event-bookie/internal/handlers"
    "github.com/rowjay007/event-bookie/internal/middleware"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    // _ "github.com/rowjay007/event-bookie/docs" 
)

// NewRouter creates a new router instance with defined routes
func NewRouter(
    // eventHandler *handlers.EventHandler,
    // userHandler *handlers.UserHandler,
    // bookingHandler *handlers.BookingHandler,
    // venueHandler *handlers.VenueHandler,
    // paymentHandler *handlers.PaymentHandler,
    // organizerHandler *handlers.OrganizerHandler,
) *gin.Engine {
    r := gin.Default()

    // Middleware
    r.Use(middleware.LoggerMiddleware())

    // Routes
    // v1 := r.Group("/api/v1")
    // {
    //     // Event routes
    //     eventGroup := v1.Group("/events")
    //     {
    //         eventGroup.GET("/", eventHandler.GetAllEvents)
    //         eventGroup.POST("/", eventHandler.CreateEvent)
    //         eventGroup.GET("/:id", eventHandler.GetEventByID)
    //         // Add more event routes as needed
    //     }

    //     // User routes
    //     userGroup := v1.Group("/users")
    //     {
    //         userGroup.GET("/", userHandler.GetAllUsers)
    //         userGroup.POST("/", userHandler.CreateUser)
    //         userGroup.GET("/:id", userHandler.GetUserByID)
    //         // Add more user routes as needed
    //     }

    //     // Booking routes
    //     bookingGroup := v1.Group("/bookings")
    //     {
    //         bookingGroup.GET("/", bookingHandler.GetAllBookings)
    //         bookingGroup.POST("/", bookingHandler.CreateBooking)
    //         bookingGroup.GET("/:id", bookingHandler.GetBookingByID)
    //         // Add more booking routes as needed
    //     }

    //     // Venue routes
    //     venueGroup := v1.Group("/venues")
    //     {
    //         venueGroup.GET("/", venueHandler.GetAllVenues)
    //         venueGroup.POST("/", venueHandler.CreateVenue)
    //         venueGroup.GET("/:id", venueHandler.GetVenueByID)
    //         // Add more venue routes as needed
    //     }

    //     // Payment routes
    //     paymentGroup := v1.Group("/payments")
    //     {
    //         paymentGroup.GET("/", paymentHandler.GetAllPayments)
    //         paymentGroup.POST("/", paymentHandler.CreatePayment)
    //         paymentGroup.GET("/:id", paymentHandler.GetPaymentByID)
    //         // Add more payment routes as needed
    //     }

    //     // Organizer routes
    //     organizerGroup := v1.Group("/organizers")
    //     {
    //         organizerGroup.GET("/", organizerHandler.GetAllOrganizers)
    //         organizerGroup.POST("/", organizerHandler.CreateOrganizer)
    //         organizerGroup.GET("/:id", organizerHandler.GetOrganizerByID)
    //         // Add more organizer routes as needed
    //     }
    // }

    // Swagger routes
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return r
}
