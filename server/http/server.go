package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/Noah-Huppert/golog"
)

// Server responds to HTTP requests
type Server struct {
	// ctx is the context used to manage server execution
	ctx context.Context

	// cfg is the application configuration
	cfg config.Config

	// logger is used to record run information
	logger golog.Logger
}

// NewServer creates a Server
func NewServer(ctx context.Context, cfg config.Config, logger golog.Logger) Server {
	return Server{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger,
	}
}

// Serve brings up an HTTP server to serve requests
func (s Server) Serve() error {
	// Load routes
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	// Create HTTP server
	httpAddr := fmt.Sprintf(":%d", s.cfg.Port)

	httpServer := http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	// Stop server when context is canceled
	shutdownChan := make(chan error)

	go func() {
		<-s.ctx.Done()

		s.logger.Info("shutting down HTTP server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			shutdownChan <- fmt.Errorf("error shutting down HTTP server: %s", err.Error())
		}

		close(shutdownChan)
	}()

	// Run HTTP server
	s.logger.Infof("starting HTTP server on %s", httpAddr)

	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("error running HTTP server: %s", err.Error())
	}

	return <-shutdownChan
}
