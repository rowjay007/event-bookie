package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq" // PostgreSQL driver
    "github.com/rowjay007/event-bookie/config"
)

// NewDB creates a new database connection
func NewDB(config *config.Config) (*sql.DB, error) {
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

    db, err := sql.Open("postgres", connectionString)
    if err != nil {
        return nil, err
    }

    // Check if the database connection is successful
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    // Apply migrations
    if err := ApplyMigrations(db); err != nil {
        return nil, err
    }

    return db, nil
}

// ApplyMigrations applies migrations to the database
func ApplyMigrations(db *sql.DB) error {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    driver := "postgres"
    migrationsPath := "file://migrations"

    if _, exists := os.LookupEnv("DATABASE_URL"); exists {
        driver = "postgres"
        migrationsPath = os.Getenv("MIGRATIONS_PATH")
    }

    // Initialize the migration instance
    m, err := migrate.New(
        migrationsPath,
        fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
            driver,
            os.Getenv("DB_USER"),
            os.Getenv("DB_PASSWORD"),
            os.Getenv("DB_HOST"),
            os.Getenv("DB_PORT"),
            os.Getenv("DB_NAME")),
    )
    if err != nil {
        return err
    }

    // Apply all available migrations
    err = m.Up()
    if err != nil && err != migrate.ErrNoChange {
        return err
    }

    return nil
}
