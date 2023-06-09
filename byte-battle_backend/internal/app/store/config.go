package store

import (
	"byte-battle_backend/pkg/loggers"
	"os"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string
}

func NewConfig() *Config {
	dbHost := os.Getenv("POSTGRES_HOST")
	if dbHost == "" {
		loggers.VariableNotFound("POSTGRES_HOST")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	if dbPort == "" {
		loggers.VariableNotFound("POSTGRES_PORT")
	}

	dbName := os.Getenv("POSTGRES_NAME")
	if dbName == "" {
		loggers.VariableNotFound("POSTGRES_NAME")
	}

	dbUser := os.Getenv("POSTGRES_USER")
	if dbUser == "" {
		loggers.VariableNotFound("POSTGRES_USER")
	}

	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	if dbPassword == "" {
		loggers.VariableNotFound("POSTGRES_PASSWORD")
	}

	dbSSLmode := os.Getenv("POSTGRES_SSLMODE")
	if dbSSLmode == "" {
		loggers.VariableNotFound("POSTGRES_SSLMODE")
	}

	return &Config{
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseName:     dbName,
		DatabaseUser:     dbUser,
		DatabasePassword: dbPassword,
		DatabaseSSLMode:  dbSSLmode,
	}
}
