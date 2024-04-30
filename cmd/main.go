// cmd/main.go

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/database"
	// "github.com/rowjay007/event-bookie/internal/routes"
)

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Initialize database connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Register routes
	// routes.RegisterRoutes(r)

	// Start server
	port := "8080"
	log.Printf("Server started on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
