package api

import (
	"net/http"

	// "github.com/rowjay007/event-bookie/internal/api/handlers"
	"github.com/rowjay007/event-bookie/internal/api/middleware"

	"github.com/gorilla/mux"
)

// NewRouter returns a new router instance
func NewRouter() http.Handler {
    // Create a new router
    r := mux.NewRouter()

    // API routes
    api := r.PathPrefix("/api").Subrouter()

    // Event routes
    // eventRouter := api.PathPrefix("/events").Subrouter()
    // eventRouter.HandleFunc("/", handlers.GetEvents).Methods("GET")
    // eventRouter.HandleFunc("/{id}", handlers.GetEvent).Methods("GET")
    // // Add more event routes as needed

    // // User routes
    // userRouter := api.PathPrefix("/users").Subrouter()
    // userRouter.HandleFunc("/", handlers.GetUsers).Methods("GET")
    // userRouter.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
    // Add more user routes as needed

    // Apply middleware
    api.Use(middleware.LoggingMiddleware)
    api.Use(middleware.ValidationMiddleware)
    // Add more middleware as needed

    return r
}
