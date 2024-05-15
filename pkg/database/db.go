package database

import (
	"database/sql"
	"fmt"

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

    return db, nil
}
