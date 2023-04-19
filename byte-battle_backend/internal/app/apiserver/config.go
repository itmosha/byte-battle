package apiserver

import (
	"byte-battle_backend/internal/app/store"
	"byte-battle_backend/pkg/loggers"
	"fmt"
	"net/http"
	"os"
)

type Config struct {
	BackendPort string
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

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func configHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/api/register/", s.handleRegister()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/api/login/", s.handleLogin()).Methods("POST", "OPTIONS")
}
