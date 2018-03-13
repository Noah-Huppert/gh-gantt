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
)

// IssuesPath is the base path issue handlers are registered at.
const IssuesPath string = "/api/issues"

// Issues implements HTTP handlers for GitHub issues
type Issues struct {
	// BasePath is the location issue handlers will be registered at
	BasePath string

	// ctx is the context.Context used to control GitHub API request
	// execution.
	ctx context.Context

	// cfg is the application configuration.
	cfg *config.Config

	// ghClient is a GitHub API client.
	ghClient *github.Client

	// redisClient is a redis cache client
	redisClient *cache.Codec
}

// NewIssues creates a new Issues instance with the default BasePath.
func NewIssues(ctx context.Context, cfg *config.Config, ghClient *github.Client,
	redisClient *cache.Codec) Issues {
	return Issues{
		BasePath:    IssuesPath,
		ctx:         ctx,
		cfg:         cfg,
		ghClient:    ghClient,
		redisClient: redisClient,
	}
}

// Register implements Registerable.Register
func (i Issues) Register(mux *http.ServeMux) {
	mux.HandleFunc(i.BasePath, i.Get)
}

// Get retrieves all GitHub issues.
func (i Issues) Get(w http.ResponseWriter, r *http.Request) {
	// Get repo
	repo, err := gh.RetrieveRepo(i.ctx, i.cfg, i.ghClient, i.redisClient)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error retrieving GitHub repo: %s",
			err.Error()))
		return
	}

	// Get issues
	issues, err := gh.RetrieveIssues(i.ctx, i.cfg, i.ghClient, i.redisClient)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error retrieving GitHub issues: %s",
			err.Error()))
		return
	}

	// Get issue dependencies
	depIssues := []zenhub.DepIssue{}
	for _, issue := range issues {
		deps, err := zenhub.RetrieveDeps(i.ctx, i.cfg, i.redisClient,
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
