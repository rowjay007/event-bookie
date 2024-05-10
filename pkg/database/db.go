// pkg/database/database.go

package database

import (
    "database/sql"
    "fmt"
    "github.com/rowjay007/event-bookie/config"
    _ "github.com/lib/pq"
)

// NewDB creates a new database connection using configuration values
func NewDB(cfg *config.Config) (*sql.DB, error) {
    dbURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
        cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
    
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %v", err)
    }
    
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %v", err)
    }
    
    return db, nil
}
