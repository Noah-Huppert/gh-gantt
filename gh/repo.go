package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

	"github.com/Noah-Huppert/gh-gantt/config"
)

// RepoCacheKey is the key to the store repository in cache under
const RepoCacheKey string = "github.repo"

// RetrieveRepo returns the repository specified in the configuration. An
// error is returned if one occurs.
func RetrieveRepo(ctx context.Context, cfg *config.Config,
	ghClient *github.Client) (*github.Repository, error) {

	// Make request
	repo, _, err := ghClient.Repositories.Get(ctx, cfg.GitHub.RepoOwner,
		cfg.GitHub.RepoName)

	if err != nil {
		return nil, fmt.Errorf("error retrieving repository: %s",
			err.Error())
	}

	return repo, nil
}
