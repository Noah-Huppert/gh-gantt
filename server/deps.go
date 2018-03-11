package server

import (
	"context"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/zenhub"
	"github.com/google/go-github/github"
)

// DepsPath is the base path dependency handlers are registered at.
const DepsPath string = "/api/dependencies"

// Deps implements HTTP handlers for dependencies between GitHub issues.
type Deps struct {
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

// NewDeps creates a new Deps instance with the default BasePath.
func NewDeps(ctx context.Context, cfg *config.Config, ghClient *github.Client) Deps {
	return Deps{
		BasePath: IssuesPath,
		ctx:      ctx,
		cfg:      cfg,
		ghClient: ghClient,
	}
}

// Register implements Registerable.Register
func (d Deps) Register(mux *http.ServeMux) {
	mux.HandleFunc(d.BasePath, d.Get)
}

// Get retrieves all GitHub issue dependencies.
func (d Deps) Get(w http.ResponseWriter, r *http.Request) {
	// Get repo id
	repo, err := gh.RetriveRepo(d.ctx, d.cfg, d.ghClient)
	if err != nil {
		WriteErr(w, 200, fmt.Errorf("error retrieving repository "+
			"information: %s", err))
		return
	}

	// Get dependencies
	deps, err := zenhub.RetriveDeps(d.ctx, d.cfg, d.ghClient)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error retrieving issue "+
			"dependencies: %s", err))
		return
	}

	resp := map[string]interface{}{
		"dependencies": deps,
	}
	WriteJSON(w, 200, resp)
}
