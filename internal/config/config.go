package config

import (
	"log"
	"os"
)

// Config holds application runtime configuration sourced from environment variables.
type Config struct {
	AppPort       string
	DatabaseURL   string
	ServerKey     string
	SessionSecret string
}

// Load reads environment variables and fills Config with sane defaults.
func Load() Config {
	cfg := Config{
		AppPort:       getEnv("APP_PORT", "8080"),
		DatabaseURL:   getEnv("DATABASE_URL", "postgres://pns:pns@localhost:5432/pns?sslmode=disable"),
		ServerKey:     os.Getenv("SERVER_KEY"),
		SessionSecret: getEnv("SESSION_SECRET", "dev-secret-change-me"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	if cfg.ServerKey == "" {
		log.Fatal("SERVER_KEY must be set to secure device API access")
	}
	if cfg.SessionSecret == "" {
		log.Fatal("SESSION_SECRET must be set for web sessions")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

