package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/zenhub"
	"github.com/google/go-github/github"

	"github.com/go-redis/cache"
	"github.com/gorilla/mux"
)

// IssuesPath is the base path issue handlers are registered at.
const IssuesPath string = "/api/issues"

// IssuesEndpoint implements HTTP handlers for GitHub issues
type IssuesEndpoint struct {
	// BasePath is the location issue handlers will be registered at
	BasePath string

	// ctx is the context.Context used to control GitHub API request
	// execution.
	ctx context.Context

	// cfg is the application configuration.
	cfg *config.Config

	// ghClient is a GitHub API client.
	ghClient *github.Client

	// redisCache is a redis cache client
	redisCache *cache.Codec
}

// NewIssuesEndpoint creates a new IssuesEndpoint instance with the default BasePath.
func NewIssuesEndpoint(ctx context.Context, cfg *config.Config, ghClient *github.Client,
	redisCache *cache.Codec) IssuesEndpoint {
	return IssuesEndpoint{
		BasePath:    IssuesPath,
		ctx:         ctx,
		cfg:         cfg,
		ghClient:    ghClient,
		redisCache: redisCache,
	}
}

// Register implements Registerable.Register
func (i IssuesEndpoint) Register(router *mux.Router) {
	router.HandleFunc(i.BasePath, i.Get).Methods("GET")
}

// Get retrieves all GitHub issues.
func (i IssuesEndpoint) Get(w http.ResponseWriter, r *http.Request) {
	// Get repo
	repo, err := gh.RetrieveRepo(i.ctx, i.cfg, i.ghClient, i.redisCache)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error retrieving GitHub repo: %s",
			err.Error()))
		return
	}

	// Get issues
	issues, err := gh.RetrieveIssues(i.ctx, i.cfg, i.ghClient, i.redisCache)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error retrieving GitHub issues: %s",
			err.Error()))
		return
	}

	// Get issue dependencies
	depIssues := []zenhub.DepIssue{}
	for _, issue := range issues {
		deps, err := zenhub.RetrieveDeps(i.ctx, i.cfg, i.redisCache,
			*repo.ID, *issue.Number)

		if err != nil {
			WriteErr(w, 500, fmt.Errorf("error retrieving "+
				"issue dependencies, issue #: %d, err: %s",
				issue.Number, err.Error()))
			return
		}

		depIss := zenhub.NewDepIssue(*issue, deps)
		depIssues = append(depIssues, depIss)
	}

	// Respond
	resp := map[string]interface{}{}
	resp["issues"] = depIssues

	WriteJSON(w, 200, resp)
}
