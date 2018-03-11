package server

import (
	"context"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/google/go-github/github"
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
}

// NewIssues creates a new Issues instance with the default BasePath.
func NewIssues(ctx context.Context, cfg *config.Config, ghClient *github.Client) Issues {
	return Issues{
		BasePath: IssuesPath,
		ctx:      ctx,
		cfg:      cfg,
		ghClient: ghClient,
	}
}

// Register implements Registerable.Register
func (i Issues) Register(mux *http.ServeMux) {
	mux.HandleFunc(i.BasePath, i.Get)
}

// Get retrieves all GitHub issues.
func (i Issues) Get(w http.ResponseWriter, r *http.Request) {
	// Get issues
	issues, err := gh.RetrieveIssues(i.ctx, i.cfg, i.ghClient)
	if err != nil {
		WriteErr(w, 500, err)
		return
	}

	resp := map[string]interface{}{
		"issues": issues,
	}
	WriteJSON(w, 200, resp)
}
