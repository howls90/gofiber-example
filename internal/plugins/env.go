package plugins

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Check if environment vars are present
func CheckEnv() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal("Error loading .env file")
    }

	if os.Getenv("DB_DNS") == "" {
		log.Fatal("env var DB_DNS is not present")
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("env var JWT_SECRET is not present")
	}

	if os.Getenv("REDIS_ADDR") == "" {
		log.Fatal("env var REDIS_ADDR is not present")
	}
}