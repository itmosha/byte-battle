package main

import (
	"byte-battle_backend/internal/app/apiserver"
	"byte-battle_backend/pkg/loggers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		loggers.LoadEnvFailure()
	}

	config := apiserver.NewConfig()

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		loggers.StartServerFailure(err)
	}
}
