package api

import (
	"context"

	"github.com/Noah-Huppert/gh-gantt/server/api/auth"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
)

// APIHandlers sets up the handlers for API endpoints
type APIHandlers struct {
	// ctx is the application context
	ctx context.Context

	// logger is used to print debug information in API endpoints
	logger golog.Logger

	// cfg is the application configuration
	cfg config.Config
}

// NewAPIHandlers creates a new APIHandlers
func NewAPIHandlers(ctx context.Context, logger golog.Logger, cfg config.Config) APIHandlers {
	return APIHandlers{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

// SetupRouter registers API routes with the provided router
func (a APIHandlers) SetupRouter(router *mux.Router) {
	healthHandler := resp.WrapResponderHandler(NewHealthCheckHandler())
	router.Handle("/healthz", healthHandler).Methods("GET")

	// Authentication
	authLogger := a.logger.GetChild("auth")

	// ... Login
	authLoginLogger := authLogger.GetChild("login")
	authLoginHandler := resp.WrapResponderHandler(auth.NewAuthLoginHandler(authLoginLogger, a.cfg))
	router.Handle("/auth/login", authLoginHandler).Methods("GET")

	// ... Exchange
	authExchangeLogger := authLogger.GetChild("exchange")
	authExchangeHandler := resp.WrapResponderHandler(auth.NewAuthExchangeHandler(a.ctx, authExchangeLogger, a.cfg))
	router.Handle("/auth/exchange", authExchangeHandler).Methods("POST")

	// ... ZenHub Append
	authZenHubAppendLogger := authLogger.GetChild("zenhub.append")
	authZenHubAppendHandler := resp.WrapResponderHandler(auth.NewZenHubAppendHandler(authZenHubAppendLogger, a.cfg))
	router.Handle("/auth/zenhub", authZenHubAppendHandler).Methods("POST")

	// Issues
	issuesLogger := a.logger.GetChild("issues")
	issuesHandler := resp.WrapResponderHandler(NewIssuesHandler(a.ctx, issuesLogger, a.cfg))
	router.Handle("/issues", issuesHandler).Methods("GET")

	// Repositories
	reposLogger := a.logger.GetChild("repos")
	reposHandler := resp.WrapResponderHandler(NewReposHandler(a.ctx, reposLogger, a.cfg))
	router.Handle("/repositories", reposHandler).Methods("GET")
}
