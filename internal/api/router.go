package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router for the API endpoints
func NewRouter() *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()

	// Define your API routes here
	router.HandleFunc("/events", handleGetEvents).Methods("GET")

	return router
}

// handleGetEvents is the handler function for GET /api/v1/events
func handleGetEvents(w http.ResponseWriter, r *http.Request) {
	// Your logic to handle GET /api/v1/events
}
