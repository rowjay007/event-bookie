package database

import (
    "database/sql"
    "fmt"

    _ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Connect connects to the database using the provided DSN
func Connect(dsn string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dsn)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }

    return db, nil
}
