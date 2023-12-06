package server

import (
	"net/http"

	"github.com/zoumas/aggrss/service/internal/env"
)

type Server struct {
	Env *env.Env
}

func New(env *env.Env) *Server {
	return &Server{
		Env: env,
	}
}

func (s *Server) ListenAndServe() error {
	srv := &http.Server{
		Addr:    ":" + s.Env.Port,
		Handler: ConfiguredRouter(),
	}
	return srv.ListenAndServe()
}

func healthCheck(w http.ResponseWriter, _ *http.Request) {
	type statusResponse struct {
		Status string `json:"status"`
	}

	respondWithJSON(
		w,
		http.StatusOK,
		statusResponse{http.StatusText(http.StatusOK)},
	)
}

func errCheck(w http.ResponseWriter, _ *http.Request) {
	respondWithError(
		w,
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
	)
}

func panicCheck(w http.ResponseWriter, _ *http.Request) {
	panic("server recovers from panics")
}
