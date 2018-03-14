package gh

import (
	"context"
	"fmt"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/go-redis/cache"
	"github.com/google/go-github/github"
)

// IssuesCacheKey is the key all GitHub issues will be stored in the cache
const IssuesCacheKey string = "github.issues"

// RetrieveIssues returns all open issues in a GitHub repository
func RetrieveIssues(ctx context.Context, cfg *config.Config,
	ghClient *github.Client, redisCache *cache.Codec) ([]*github.Issue, error) {

	// Check if issue exists
	var issues []*github.Issue

	if err := redisCache.Get(IssuesCacheKey, &issues); (err != nil) && (err != cache.ErrCacheMiss) {
		return nil, fmt.Errorf("error retrieving all GitHub issues "+
			"from cache: %s", err.Error())
	} else if err != cache.ErrCacheMiss {
		// Cached
		return issues, nil
	}

	// Get all issues for repo
	issues, _, err := ghClient.Issues.ListByRepo(ctx, cfg.GitHub.RepoOwner,
		cfg.GitHub.RepoName, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing GitHub issues for "+
			"repository: %s/%s, err: %s", cfg.GitHub.RepoOwner,
			cfg.GitHub.RepoName, err.Error())
	}

	// Save in cache
	if err = redisCache.Set(&cache.Item{
		Key:        IssuesCacheKey,
		Object:     issues,
		Expiration: 0,
	}); err != nil {

		return nil, fmt.Errorf("error saving results to cache: %s",
			err.Error())
	}

	return issues, nil
}
