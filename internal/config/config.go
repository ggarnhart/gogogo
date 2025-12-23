package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	// AllowedOrigins   string
	// StripeSecretKey  string
	// StripeWebhookKey string
	// RedisURL         string
	Environment string
}

func Load() *Config {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Printf("No .env file loaded. Config failing to setup.")
		}
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
