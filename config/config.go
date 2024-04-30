package config

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v2"
)

type Config struct {
    Server   ServerConfig   `yaml:"server"`
    Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
    Host string `yaml:"host"`
    Port string `yaml:"port"`
}

type DatabaseConfig struct {
    Dialect          string `yaml:"dialect"`
    ConnectionString string `yaml:"connection_string"`
}

func LoadConfig() (*Config, error) {
    env := "default"
    if value, ok := os.LookupEnv("APP_ENV"); ok {
        env = value
    }

    filename := fmt.Sprintf("config/%s.yaml", env)
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var cfg Config
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        return nil, err
    }

    return &cfg, nil
}