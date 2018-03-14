package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/google/go-github/github"

	"github.com/go-redis/cache"
)

// TODO: Make Server struct
// TODO: Make static files handler

// Server wraps the HTTP server functionality
type Server struct {
	ctx         context.Context
	cfg         *config.Config
	ghClient    *github.Client
	redisClient *cache.Codec
}

// NewServer creates a new Server instance.
func NewServer(ctx context.Context, cfg *config.Config,
	ghClient *github.Client, redisClient *cache.Codec) Server {

	return Server{
		ctx:         ctx,
		cfg:         cfg,
		ghClient:    ghClient,
		redisClient: redisClient,
	}
}

// Registerables returns a slice of Registerable interface to be registered
// with a http.ServeMux.
func (s Server) Registerables() []Registerable {

	return []Registerable{
		NewStaticFiles(),
		NewIssuesEndpoint(s.ctx, s.cfg, s.ghClient, s.redisClient),
		NewPurgeEndpoint(),
	}
}

// Routes returns a http.ServeMux with all the server route handlers.
func (s Server) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	regs := s.Registerables()
	for _, reg := range regs {
		reg.Register(mux)
	}

	return mux
}

func (s Server) Start() error {
	mux := s.Routes()

	portStr := fmt.Sprintf(":%d", s.cfg.HTTP.Port)

	err := http.ListenAndServe(portStr, mux)
	if err != nil {
		return fmt.Errorf("error starting server: %s", err.Error())
	}

	return nil
}
