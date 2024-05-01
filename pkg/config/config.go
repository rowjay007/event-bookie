package config

import (
    "github.com/spf13/viper"
)

// Config represents the application configuration structure
type Config struct {
    Database DatabaseConfig
}

// DatabaseConfig represents the configuration for the database
type DatabaseConfig struct {
    Driver   string
    Host     string
    Port     string
    Name     string
    User     string
    Password string
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(configFile string) (*Config, error) {
    var config Config

    // Set default values
    viper.SetDefault("database.driver", "sqlite3")
    viper.SetDefault("database.host", "")
    viper.SetDefault("database.port", "")
    viper.SetDefault("database.name", "event_bookie.db")
    viper.SetDefault("database.user", "")
    viper.SetDefault("database.password", "")

    // Read config from file
    viper.SetConfigFile(configFile)
    err := viper.ReadInConfig()
    if err != nil {
        return nil, err
    }

    // Unmarshal config into struct
    err = viper.Unmarshal(&config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}
