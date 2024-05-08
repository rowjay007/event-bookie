package main

import (
	"context" // Added import for database/sql
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/event-bookie/internal/api"
	"github.com/rowjay007/event-bookie/pkg/config"
	"github.com/rowjay007/event-bookie/pkg/database"
)

func main() {
    // Load application configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }

    // Initialize database
    db, err := database.InitDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("failed to initialize database: %v", err)
    }
    defer db.Close()

    // Initialize Gin router
    router := gin.Default()

    // Register API routes
    api.RegisterRoutes(router, db)

    // Start HTTP server
    server := &http.Server{
        Addr:    fmt.Sprintf(":%d", cfg.Port),
        Handler: router,
    }
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("failed to start server: %v", err)
        }
    }()
    log.Printf("Server started on port %d", cfg.Port)

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("Shutting down server...")

    // Create a context with timeout for graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Shut down server gracefully with a timeout
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("failed to gracefully shut down server: %v", err)
    }
    log.Println("Server stopped")
}
