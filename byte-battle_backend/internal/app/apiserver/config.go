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

func logNotSpecifiedError(varName string) {
	log.Fatalf("SERVER CONFIG ERROR: %s variable was not specified in .env file", varName)
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	backendPort, logLevel := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT")), os.Getenv("LOG_LEVEL")

	if backendPort == ":" {
		logNotSpecifiedError("BACKEND_PORT")
	}
	if logLevel == "" {
		logNotSpecifiedError("LOG_LEVEL")
	}

	return &Config{
		BackendPort: backendPort,
		LogLevel:    logLevel,
		Store:       store.NewConfig(),
	}
}
