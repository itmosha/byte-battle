package apiserver

import (
	"fmt"
	"log"
	"os"

	"github.com/Mihalych32/byte-battle/tree/main/byte-battle_backend/internal/app/store"
	"github.com/joho/godotenv"
)

type Config struct {
	BackendPort string
	LogLevel    string
	Store       *store.Config
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backend_port, log_level := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")), os.Getenv("LOG_LEVEL")

	if backend_port == ":" {
		log.Fatal("ERROR: BACKEND_PORT was not specified in .env file")
	}
	if log_level == "" {
		log.Fatal("ERROR: LOG_LEVEL was not specified in .env file")
	}

	return &Config{
		BackendPort: backend_port,
		LogLevel:    log_level,
		Store:       store.NewConfig(),
	}
}
