package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/google/go-github/github"

	"github.com/go-redis/cache"
	"github.com/go-redis/redis"

	"github.com/gorilla/mux"
)

// TODO: Make Server struct
// TODO: Make static files handler

// Server wraps the HTTP server functionality
type Server struct {
	ctx         context.Context
	cfg         *config.Config
	ghClient    *github.Client
	redisClient *redis.Client
	redisCache  *cache.Codec
}

// NewServer creates a new Server instance.
func NewServer(ctx context.Context, cfg *config.Config,
	ghClient *github.Client, redisClient *redis.Client,
	redisCache *cache.Codec) Server {

	return Server{
		ctx:         ctx,
		cfg:         cfg,
		ghClient:    ghClient,
		redisClient: redisClient,
		redisCache:  redisCache,
	}
}

// Registerables returns a slice of Registerable interface to be registered
// with a http.ServeMux.
func (s Server) Registerables() []Registerable {

	return []Registerable{
		NewNotFoundHandler(),
		NewIssuesEndpoint(s.ctx, s.cfg, s.ghClient, s.redisClient, s.redisCache),
		NewPurgeEndpoint(s.redisClient, s.redisCache),
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
