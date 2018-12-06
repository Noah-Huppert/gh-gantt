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

// IssuesHandler implements resp.ResponderHandler by returning a list of GitHub issues for a repository
type IssuesHandler struct {
	// ctx is application context
	ctx context.Context

	// logger prints debug information
	logger golog.Logger

	// cfg is application configuration
	cfg config.Config
}

// IssuesRequest is the format of a issues request
type IssuesRequest struct {
	// RepositoryOwner is the repository's owner
	RepositoryOwner string `json:"repository_owner" validate:"nonzero"`

	// RepositoryName is the repository's name
	RepositoryName string `json:"repository_name" validate:"nonzero"`
}

// NewIssuesHandler creates a new IssuesHandler
func NewIssuesHandler(ctx context.Context, logger golog.Logger, cfg config.Config) IssuesHandler {
	return IssuesHandler{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h IssuesHandler) Handle(r *http.Request) resp.Responder {
	// Check auth token
	authToken, errResp := req.CheckAuthToken(h.logger, h.cfg, r)
	if errResp != nil {
		return errResp
	}

	// Get query parameters
	// TODO: Make helper
	var request IssuesRequest

	repoOwner := r.URL.Query().Get("repository_owner")
	if len(repoOwner) == 0 {
		return resp.NewStrErrorResponder(h.logger, http.StatusBadRequest,
			"repository_owner query parameter must be provided", "")
	}

	repoName := r.URL.Query().Get("repository_name")
	if len(repoOwner) == 0 {
		return resp.NewStrErrorResponder(h.logger, http.StatusBadRequest,
			"repository_name query parameter must be provided", "")
	}

	request.RepositoryOwner = repoOwner
	request.RepositoryName = repoName

	// Get GitHub issues
	client := github.NewUserClient(h.ctx, authToken.GitHubAuthToken)

	issues, _, err := client.Issues.ListByRepo(h.ctx, request.RepositoryOwner, request.RepositoryName, nil)
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error retrieving GitHub issues", err.Error())
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"issues": issues,
	}, http.StatusOK)
}
