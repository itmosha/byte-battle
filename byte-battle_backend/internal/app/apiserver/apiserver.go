package apiserver

import (
	"byte-battle_backend/internal/app/store"
	"byte-battle_backend/pkg/loggers"
	"io"
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

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/api", s.handleApiList())
}

func (s *APIserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIserver) handleApiList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "API root")
	}
}
