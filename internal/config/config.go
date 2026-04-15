package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Add this
)

type Config struct {
	Port        string
	DatabaseURL string
	Env         string // "development" or "production"
}

func Load() *Config {
	// This looks for a .env file in the root and loads it into the environment
	// We ignore the error so the app still works if the file is missing (like in Production)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}

	return &Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Env:         getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
