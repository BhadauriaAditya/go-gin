package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Config struct {
	DSN string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config.DSN = os.Getenv("DB_DSN")
}
