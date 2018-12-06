package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/Noah-Huppert/gh-gantt/server/auth"
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
	// TODO: Turn into middleware
	authorization := r.Header.Get("Authorization")
	if len(authorization) == 0 {
		return resp.NewStrErrorResponder(h.logger, http.StatusUnauthorized,
			"no authentication token provided", "authorization header empty")
	}

	authorizationParts := strings.Split(authorization, " ")
	if len(authorizationParts) != 2 {
		return resp.NewStrErrorResponder(h.logger, http.StatusUnauthorized,
			"authorization header not in correct format, must be: token <TOKEN>", "")
	}

	authTokenStr := authorizationParts[1]

	authToken := &auth.AuthToken{}
	err := authToken.Decode(authTokenStr, h.cfg.SigningSecret)

	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error parsing provided API authentication token", err.Error())
	}

	// Decode body
	var request IssuesRequest

	errResp := req.DecodeValidatedJSON(h.logger, r, &request)
	if errResp != nil {
		return errResp
	}

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
