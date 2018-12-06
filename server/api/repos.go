package api

import (
	"context"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/auth/github"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/req"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
)

// ReposHandler implements resp.ResponderHandler by returning a list of GitHub repository names for a user
type ReposHandler struct {
	// ctx is application context
	ctx context.Context

	// logger prints debug information
	logger golog.Logger

	// cfg is application configuration
	cfg config.Config
}

// NewReposHandler creates a new ReposHandler
func NewReposHandler(ctx context.Context, logger golog.Logger, cfg config.Config) ReposHandler {
	return ReposHandler{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h ReposHandler) Handle(r *http.Request) resp.Responder {
	// Check auth token
	authToken, errResp := req.CheckAuthToken(h.logger, h.cfg, r)
	if errResp != nil {
		return errResp
	}

	// Get GitHub repos
	client := github.NewUserClient(h.ctx, authToken.GitHubAuthToken)

	repos, _, err := client.Repositories.ListAll(h.ctx, nil)
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error retrieving repositories from GitHub API", err.Error())
	}

	// Group repositories for response
	reposResp := map[string][]string{}

	for _, repo := range repos {
		// Check if org exists in resp
		org := *(repo.Organization.Login)

		if _, ok := reposResp[org]; !ok {
			reposResp[org] = []string{}
		}

		// Add to resp
		reposResp[org] = append(reposResp[org], *(repo.Name))
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"repositories": reposResp,
	}, http.StatusOK)
}
