package database

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

    // Apply migrations using configuration settings
    if err := ApplyMigrations(db, config); err != nil {
        return nil, err
    }

    return db, nil
}

func ApplyMigrations(db *sql.DB, config *config.Config) error {
    driver := "postgres"
    migrationsPath := "file://../migrations"

    // Initialize the migration instance
    m, err := migrate.New(
        migrationsPath,
        fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
            driver,
            config.DBUser,
            config.DBPassword,
            config.DBHost,
            config.DBPort,
            config.DBName),
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

