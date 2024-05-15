package config

import (
    "os"
)

// Config holds the configuration values for the application
type Config struct {
    Port       string
    DBUser     string
    DBPassword string
    DBHost     string
    DBPort     string
    DBName     string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
    return &Config{
        Port:       getEnv("PORT", ""),
        DBUser:     getEnv("DB_USER", ""),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBName:     getEnv("DB_NAME", ""),
    }
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
