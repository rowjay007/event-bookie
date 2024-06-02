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
// @termsOfService http://swagger.io/terms/
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

    db, gormDB, err := initializeDatabase(conf, logger)
    if err != nil {
        logger.Fatalf("Error initializing database: %v", err)
    }
    defer func() {
        if err := db.Close(); err != nil {
            logger.Errorf("Error closing the database: %v", err)
        }
    }()

    logger.Info("Connected to the database successfully")

    if err := database.ApplyMigrations(conf, logger); err != nil {
        logger.Errorf("Error applying migrations: %v", err)
    } else {
        logger.Info("Migrations applied successfully")
    }

    r := router.SetupRouter(gormDB)

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

    waitForShutdown(server, logger)
}

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
