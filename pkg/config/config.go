package config

import (
	"os"
)

// Config represents the application configuration
type Config struct {
	// Add other configuration fields as needed
	PostgresURL string
}

// LoadConfig loads the application configuration from various sources
func LoadConfig() *Config {
	return &Config{
		PostgresURL: os.Getenv("POSTGRES_URL"),
	}
}
