package apiserver

import (
	"byte-battle_backend/internal/app/store"
	"byte-battle_backend/pkg/loggers"
	"net/http"

	"github.com/gorilla/mux"
)

type APIserver struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	loggers.StartServerSuccess()

	return http.ListenAndServe(s.config.BackendPort, s.router)
}
