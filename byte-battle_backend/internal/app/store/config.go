package store

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_url := os.Getenv("POSTGRES_URL")

	if db_url == "" {
		log.Fatal("ERROR: POSTGRES_URL was not specified in .env file")
	}

	return &Config{
		DatabaseURL: db_url,
	}
}
