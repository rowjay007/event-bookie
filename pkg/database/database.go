package database

import (
    "database/sql"
    _ "github.com/lib/pq"
)

// Connect establishes a connection to the database
func Connect(url string) (*sql.DB, error) {
    db, err := sql.Open("postgres", url)
    if err != nil {
        return nil, err
    }
    // Ping the database to ensure the connection is valid
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
