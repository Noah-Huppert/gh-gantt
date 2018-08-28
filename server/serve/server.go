package serve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/api"
	"github.com/Noah-Huppert/gh-gantt/server/config"

	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// Server responds to HTTP requests
type Server struct {
	// ctx is the context used to manage server execution
	ctx context.Context

	// cfg is the application configuration
	cfg config.Config

	// db is a database connection
	db *sqlx.DB

	// logger is used to record run information
	logger golog.Logger
}

// NewServer creates a Server
func NewServer(ctx context.Context, cfg config.Config, db *sqlx.DB, logger golog.Logger) Server {
	return Server{
		ctx:    ctx,
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
}

// Serve brings up an HTTP server to serve requests
func (s Server) Serve() error {
	// Load routes
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v0").Subrouter()

	apiHandlers := api.NewAPIHandlers(s.logger, s.cfg, s.db)
	apiHandlers.SetupRouter(apiRouter)

	router.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	// Setup recovery handler
	recoveryHandler := NewRecoveryHandler(router, s.logger.GetChild("recovery"))

	// Create HTTP server
	httpAddr := fmt.Sprintf(":%d", s.cfg.Port)

	httpServer := http.Server{
		Addr:    httpAddr,
		Handler: recoveryHandler,
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
