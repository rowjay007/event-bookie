package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3" // SQLite driver
)

// InitDB initializes the database connection
func InitDB(dataSourceName string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Perform a ping to check if the database is accessible
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    log.Println("Connected to database")
    return db, nil
}
