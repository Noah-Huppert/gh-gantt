package api

import (
	"context"
	"net/http"
	"strings"

	libgithub "github.com/Noah-Huppert/gh-gantt/server/auth/github"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/req"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"github.com/google/go-github/github"
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
	client := libgithub.NewUserClient(h.ctx, authToken.GitHubAuthToken)

	// TODO: Get paginated results so all repos will be returned
	// TODO: Get all orgs user belongs to and call list for each
	repos, _, err := client.Repositories.List(h.ctx, "", &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			PerPage: 999,
		},
	})
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error retrieving repositories from GitHub API", err.Error())
	}

	// Group repositories for response
	reposResp := map[string][]string{}

	for _, repo := range repos {
		// Check if owner exists in resp
		nameParts := strings.Split(*(repo.FullName), "/")

		owner := nameParts[0]
		name := nameParts[1]

		if _, ok := reposResp[owner]; !ok {
			reposResp[owner] = []string{}
		}

		// Add to resp
		reposResp[owner] = append(reposResp[owner], name)
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"repositories": reposResp,
	}, http.StatusOK)
}
