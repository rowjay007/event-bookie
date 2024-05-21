package config

import (
    "log"
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
    config := &Config{
        Port:       getEnv("PORT", "8080"),
        DBUser:     getEnv("DB_USER", ""),
        DBPassword: getEnv("DB_PASSWORD", ""),
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBName:     getEnv("DB_NAME", ""),
    }

    // Log a warning if non-essential environment variables are missing
    checkEnv("PORT")
    checkEnv("DB_USER")
    checkEnv("DB_PASSWORD")
    checkEnv("DB_HOST")
    checkEnv("DB_PORT")
    checkEnv("DB_NAME")

    return config
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

// checkEnv logs a warning if an environment variable is not set
func checkEnv(key string) {
    if _, exists := os.LookupEnv(key); !exists {
        log.Printf("Warning: Environment variable %s is not set", key)
    }
}
