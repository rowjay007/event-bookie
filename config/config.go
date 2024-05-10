// config/config.go

package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config struct to hold configuration variables
type Config struct {
    Port       string
    Database   DatabaseConfig
    // Add more configuration variables as needed
}

// DatabaseConfig struct to hold database configuration variables
type DatabaseConfig struct {
    User     string
    Password string
    Host     string
    Port     string
    Name     string
}

func NewConfig() (*Config, error) {
    // Specify the path to the .env file
    err := godotenv.Load("../.env")
    if err != nil {
        return nil, err
    }

    // Parse configuration from environment variables
    port := os.Getenv("PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    // Initialize Config struct
    cfg := &Config{
        Port: port,
        Database: DatabaseConfig{
            User:     dbUser,
            Password: dbPassword,
            Host:     dbHost,
            Port:     dbPort,
            Name:     dbName,
        },
        // Add more configuration variables as needed
    }

    return cfg, nil
}


