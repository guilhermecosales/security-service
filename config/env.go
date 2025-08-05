package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Environment     string // TODO: validate enabled values (development, test and production)
	ApplicationName string
	ApplicationPort string
	DatabaseConfig  DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func LoadEnvironmentVariables() (*Config, error) {
	cfg := &Config{
		Environment:     getEnv("ENVIRONMENT", "production"),
		ApplicationName: getEnv("APP_NAME", "security-service"),
		ApplicationPort: getEnv("APPLICATION_PORT", "3001"),
		DatabaseConfig: DatabaseConfig{
			Host:     getEnv("DATABASE_HOST", "localhost"),
			Port:     getEnv("DATABASE_PORT", "5432"),
			Username: getEnv("DATABASE_USER", "postgres"),
			Password: getEnv("DATABASE_PASSWORD", "postgres"),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
