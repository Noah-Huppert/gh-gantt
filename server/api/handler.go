package api

import (
	"github.com/Noah-Huppert/gh-gantt/server/api/auth"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
)

// APIHandlers sets up the handlers for API endpoints
type APIHandlers struct {
	// logger is used to print debug information in API endpoints
	logger golog.Logger

	// cfg is the application configuration
	cfg config.Config
}

// NewAPIHandlers creates a new APIHandlers
func NewAPIHandlers(logger golog.Logger, cfg config.Config) APIHandlers {
	return APIHandlers{
		logger: logger,
		cfg:    cfg,
	}
}

// SetupRouter registers API routes with the provided router
func (a APIHandlers) SetupRouter(router *mux.Router) {
	healthHandler := resp.WrapResponderHandler(NewHealthCheckHandler())
	router.Handle("/healthz", healthHandler).Methods("GET")

	authLoginHandler := resp.WrapResponderHandler(auth.NewAuthLoginHandler(a.logger, a.cfg))
	router.Handle("/auth/login", authLoginHandler).Methods("GET")
}
