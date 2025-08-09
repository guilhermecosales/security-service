package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Environment     string // TODO: validate enabled values (development, test and production)
	ApplicationName string
	ApplicationPort string
	DatabaseConfig  DatabaseConfig
}

type DatabaseConfig struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
}

func LoadEnvironmentVariables() (*Config, error) {
	cfg := &Config{
		Environment:     getEnv("ENVIRONMENT", "production"),
		ApplicationName: getEnv("APP_NAME", "security-service"),
		ApplicationPort: getEnv("APPLICATION_PORT", "3001"),
		DatabaseConfig: DatabaseConfig{
			Host:         getEnv("DATABASE_HOST", "localhost"),
			Port:         getEnv("DATABASE_PORT", "5432"),
			Username:     getEnv("DATABASE_USER", "postgres"),
			Password:     getEnv("DATABASE_PASSWORD", "postgres"),
			DatabaseName: getEnv("DATABASE_NAME", "security_service"),
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Warn().Msg("Environment variable " + key + " not set, using default value " + fallback)
	return fallback
}
