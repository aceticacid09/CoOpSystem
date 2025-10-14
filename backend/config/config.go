package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config โครงสร้างหลักของการตั้งค่าแอป
type Config struct {
	AppPort     string // พอร์ตที่แอปจะฟัง
	DatabaseURL string // URL สำหรับเชื่อมต่อ Postgres
	JWTSecret   string // JWT secret key
}

// New โหลดค่าต่างๆ จาก environment variables ด้วย Viper
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
		cfg.JWTSecret = "your-secret-key" // todo: set this via environment
	}

	// สร้าง DatabaseURL
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
