package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Define a struct type
type Config struct {
	DSN string
}

// Global variable to hold loaded config
var AppConfig *Config

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		DSN: os.Getenv("DB_DSN"),
	}

	AppConfig = cfg
	return cfg
}