package config

import (
	"fmt"
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
		AppPort:   viper.GetString("APP_PORT"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}

	// Set default values
	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}
	if cfg.JWTSecret == "" {
		cfg.JWTSecret = "your-secret-key"
	}

	// Parse ALLOWED_ORIGINS
	allowedOrigins := viper.GetString("ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		cfg.AllowedOrigins = strings.Split(allowedOrigins, ",")
	} else {
		// Default for development
		cfg.AllowedOrigins = []string{"http://localhost:5173"}
	}

	// Build DatabaseURL
	cfg.DatabaseURL = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_DB"),
		viper.GetString("POSTGRES_SSLMODE"),
	)

	// Validate required fields
	if cfg.AppPort == "" {
		return nil, fmt.Errorf("APP_PORT is required")
	}

	return cfg, nil
}