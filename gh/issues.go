package gh

import (
	"context"
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/google/go-github/github"
)

// IssuesCacheKey is the key all GitHub issues will be stored in the cache
const IssuesCacheKey string = "github.issues"

// RetrieveIssues returns all open issues in a GitHub repository
func RetrieveIssues(ctx context.Context, cfg *config.Config,
	ghClient *github.Client) ([]*github.Issue, error) {

	// Get all issues for repo
	issues, _, err := ghClient.Issues.ListByRepo(ctx, cfg.GitHub.RepoOwner,
		cfg.GitHub.RepoName, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing GitHub issues for "+
			"repository: %s/%s, err: %s", cfg.GitHub.RepoOwner,
			cfg.GitHub.RepoName, err.Error())
	}

	return issues, nil
}
