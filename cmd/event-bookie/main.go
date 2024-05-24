// Package main is the entry point of the Event Bookie application.
package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/joho/godotenv"
    "github.com/rowjay007/event-bookie/config"
    "github.com/rowjay007/event-bookie/internal/router"
    "github.com/rowjay007/event-bookie/pkg/database"
    "github.com/sirupsen/logrus"
    _ "github.com/rowjay007/event-bookie/cmd/docs" 
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// @title Event Bookie API
// @version 1.0
// @description The Event Bookie API is a powerful platform designed to streamline event management and booking processes. It provides a comprehensive set of features for creating, managing, and discovering events, handling bookings, managing venues, processing payments, and facilitating organizer interactions.
//
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
//
// @host localhost:8080
// @BasePath /api/v1
func main() {
    if err := godotenv.Load("../../.env"); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    conf := config.NewConfig()
    logger := logrus.New()
    logger.SetFormatter(&logrus.JSONFormatter{})

    // Connect to the database and ensure proper closure
    db, gormDB, err := initializeDatabase(conf, logger)
    if err != nil {
        logger.Fatalf("Error initializing database: %v", err)
    }
    defer func() {
        if err := db.Close(); err != nil {
            logger.Errorf("Error closing the database: %v", err)
        }
    }()

    // Print a message to indicate successful database connection
    logger.Info("Connected to the database successfully")

    // Apply migrations
    if err := database.ApplyMigrations(conf, logger); err != nil {
        logger.Errorf("Error applying migrations: %v", err)
    } else {
        logger.Info("Migrations applied successfully")
    }

    // Initialize router with gormDB
    r := router.SetupRouter(gormDB)

    // Start HTTP server
    port := conf.Port
    logger.Infof("Server is running on port %s", port)

    server := &http.Server{
        Addr:    ":" + port,
        Handler: r,
    }

    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatalf("Could not listen on %s: %v\n", port, err)
        }
    }()

    // Graceful shutdown
    waitForShutdown(server, logger)
}

// initializeDatabase sets up the database connection and returns both *sql.DB and *gorm.DB
func initializeDatabase(conf *config.Config, logger *logrus.Logger) (*sql.DB, *gorm.DB, error) {
    db, err := database.NewDB(conf, logger)
    if err != nil {
        return nil, nil, fmt.Errorf("error connecting to the database: %w", err)
    }

    gormDB, err := gorm.Open(postgres.New(postgres.Config{
        Conn: db,
    }), &gorm.Config{})
    if err != nil {
        return nil, nil, fmt.Errorf("error converting *sql.DB to *gorm.DB: %w", err)
    }

    return db, gormDB, nil
}

// waitForShutdown waits for interrupt signals to gracefully shut down the server
func waitForShutdown(server *http.Server, logger *logrus.Logger) {
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    logger.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(ctx); err != nil {
        logger.Fatalf("Server forced to shutdown: %v", err)
    }

    logger.Println("Server gracefully stopped")
}