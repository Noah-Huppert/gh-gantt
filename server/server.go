package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/google/go-github/github"

	"github.com/gorilla/mux"
)

// Server wraps the HTTP server functionality
type Server struct {
	ctx      context.Context
	cfg      *config.Config
	ghClient *github.Client
}

// NewServer creates a new Server instance.
func NewServer(ctx context.Context, cfg *config.Config,
	ghClient *github.Client) Server {

	return Server{
		ctx:      ctx,
		cfg:      cfg,
		ghClient: ghClient,
	}
}

// Registerables returns a slice of Registerable interface to be registered
// with a http.ServeMux.
func (s Server) Registerables() []Registerable {

	return []Registerable{
		NewNotFoundHandler(),
		NewIssuesEndpoint(s.ctx, s.cfg, s.ghClient),
		NewStaticFiles(),
	}
}

// Routes returns a mux.Router with all the server route handlers.
func (s Server) Routes() *mux.Router {
	router := mux.NewRouter()

	regs := s.Registerables()
	for _, reg := range regs {
		reg.Register(router)
	}

	return router
}

func (s Server) Start() error {
	router := s.Routes()

	portStr := fmt.Sprintf(":%d", s.cfg.HTTP.Port)

	err := http.ListenAndServe(portStr, router)
	if err != nil {
		return fmt.Errorf("error starting server: %s", err.Error())
	}

	return nil
}
