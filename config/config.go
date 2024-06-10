package config

import (
    "log"
    "os"
    "strconv"
)

type Config struct {
    Port            string
    DBUser          string
    DBPassword      string
    DBHost          string
    DBPort          string
    DBName          string
    PaystackLiveKey string
    PaystackTestKey string
    FlutterwaveLiveKey string
    FlutterwaveTestKey string
    FlutterwaveEncryptionKey string
}

func NewConfig() *Config {
    config := &Config{
        Port:            getEnv("PORT", "8080"),
        DBUser:          getEnv("DB_USER", ""),
        DBPassword:      getEnv("DB_PASSWORD", ""),
        DBHost:          getEnv("DB_HOST", "localhost"),
        DBPort:          getEnv("DB_PORT", "5432"),
        DBName:          getEnv("DB_NAME", ""),
        PaystackLiveKey: getEnv("PAYSTACK_LIVE_KEY", ""),
        PaystackTestKey: getEnv("PAYSTACK_TEST_KEY", ""),
        FlutterwaveLiveKey: getEnv("FLUTTERWAVE_LIVE_KEY", ""),
        FlutterwaveTestKey: getEnv("FLUTTERWAVE_TEST_KEY", ""),
        FlutterwaveEncryptionKey: getEnv("FLUTTERWAVE_ENCRYPTION_KEY", ""),
    }

    checkEnv("PORT")
    checkEnv("DB_USER")
    checkEnv("DB_PASSWORD")
    checkEnv("DB_HOST")
    checkEnv("DB_PORT")
    checkEnv("DB_NAME")
    checkEnv("PAYSTACK_LIVE_KEY")
    checkEnv("PAYSTACK_TEST_KEY")
    checkEnv("FLUTTERWAVE_LIVE_KEY")
    checkEnv("FLUTTERWAVE_TEST_KEY")
    checkEnv("FLUTTERWAVE_ENCRYPTION_KEY")

    return config
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
    if value, exists := os.LookupEnv(key); exists {
        boolValue, err := strconv.ParseBool(value)
        if err == nil {
            return boolValue
        }
    }
    return defaultValue
}

func checkEnv(key string) {
    if _, exists := os.LookupEnv(key); !exists {
        log.Printf("Warning: Environment variable %s is not set", key)
    }
}
