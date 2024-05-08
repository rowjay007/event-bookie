package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
    DatabaseURL string `mapstructure:"DATABASE_URL"`
    Port        int    `mapstructure:"PORT"`
    // Add more configuration fields as needed
}

// LoadConfig loads the application configuration from config.yaml file
// LoadConfig loads the application configuration from config.yaml file
// LoadConfig loads the application configuration from config.yaml file
func LoadConfig() (*Config, error) {
    // Set default values for configuration fields
    viper.SetDefault("PORT", 8080)

    // Read configuration from config.yaml file
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("../../pkg/config/") // Corrected path
    err := viper.ReadInConfig()
    if err != nil {
        log.Fatalf("failed to read config file: %v", err)
    }

    // Unmarshal configuration into Config struct
    var cfg Config
    err = viper.Unmarshal(&cfg)
    if err != nil {
        log.Fatalf("failed to unmarshal config: %v", err)
    }

    return &cfg, nil
}
