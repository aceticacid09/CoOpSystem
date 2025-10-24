package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort        string
	DatabaseURL    string
	JWTSecret      string
	AllowedOrigins []string
}

func New() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := &Config{
		AppPort:   getEnv("APP_PORT", "8000"),
		JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
	}

	// Check if Railway's DATABASE_URL exists (prioritize it)
	if databaseURL := os.Getenv("DATABASE_URL"); databaseURL != "" {
		cfg.DatabaseURL = databaseURL
	} else {
		// Fallback to individual variables
		cfg.DatabaseURL = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			getEnv("POSTGRES_HOST", "localhost"),
			getEnv("POSTGRES_PORT", "5432"),
			getEnv("POSTGRES_USER", "postgres"),
			getEnv("POSTGRES_PASSWORD", ""),
			getEnv("POSTGRES_DB", "postgres"),
			getEnv("POSTGRES_SSLMODE", "disable"),
		)
	}

	// Parse ALLOWED_ORIGINS
	allowedOrigins := getEnv("ALLOWED_ORIGINS", "http://localhost:5173")
	if allowedOrigins != "" {
		cfg.AllowedOrigins = strings.Split(allowedOrigins, ",")
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}