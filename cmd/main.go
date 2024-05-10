package main

import (
    "fmt"
    "log"
    "github.com/rowjay007/event-bookie/config"
    "github.com/rowjay007/event-bookie/internal/router"
    "github.com/rowjay007/event-bookie/pkg/database"
)

func main() {
    // Load configuration
    cfg, err := config.NewConfig()
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }

    // Connect to the database
    db, err := database.NewDB(cfg)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    } else {
        log.Printf("Connected to database successfully")
    }
    defer db.Close()

    // Initialize Gin router
    r := router.NewRouter()

    // Run the server with automatic reload
    addr := fmt.Sprintf(":%s", cfg.Port)
    log.Printf("Server is running on http://localhost%s\n", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
