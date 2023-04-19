package apiserver

import (
	"byte-battle_backend/internal/app/store"
	"byte-battle_backend/pkg/loggers"
	"fmt"
	"os"
)

type Config struct {
	BackendPort string
	LogLevel    string
	Store       *store.Config
}

func NewConfig() *Config {

	backendPort := fmt.Sprintf(":%s", os.Getenv("BACKEND_PORT"))
	if backendPort == ":" {
		loggers.VariableNotFound("BACKEND_PORT")
	}

	return &Config{
		BackendPort: backendPort,
		Store:       store.NewConfig(),
	}
}
