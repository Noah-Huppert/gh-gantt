package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

	"github.com/Noah-Huppert/gh-gantt/config"
)

// RetrieveRepo returns the repository specified. An error is returned if
// one occurs.
func RetrieveRepo(ctx context.Context, cfg *config.Config,
	ghClient *github.Client) (*github.Repository, error) {

	// Make request
	repo, _, err := ghClient.Repositories.Get(ctx, cfg.GitHub.RepoOwner,
		cfg.GitHub.RepoName)

	if err != nil {
		return nil, fmt.Errorf("erro retrieving repository: %s",
			err.Error())
	}

	return repo, nil
}
