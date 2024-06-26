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
	_ "github.com/rowjay007/event-bookie/cmd/docs"
	"github.com/rowjay007/event-bookie/config"
	"github.com/rowjay007/event-bookie/internal/router"
	"github.com/rowjay007/event-bookie/pkg/database"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var limiter = rate.NewLimiter(rate.Limit(100), 60) 

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
        Handler: rateLimitMiddleware(r), 
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

func rateLimitMiddleware(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        handler.ServeHTTP(w, r)
    })
}
