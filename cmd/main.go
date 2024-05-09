package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rowjay007/event-bookie/internal/api"
	"github.com/rowjay007/event-bookie/pkg/database"
)

func main() {
	// Read PostgreSQL URL from environment variable
	pgURL := os.Getenv("POSTGRES_URL")
	if pgURL == "" {
		log.Fatal("POSTGRES_URL environment variable is not set")
	}

	// Initialize database connection
	db, err := database.Connect(pgURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a new API router
	router := api.NewRouter()

	// Start API server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			log.Fatalf("API server error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	log.Println("Server gracefully stopped.")
}
