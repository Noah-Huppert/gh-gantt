package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/go-redis/cache"
)

// RepoCacheKey is the key to the store repository in cache under
const RepoCacheKey string = "gh.repo"

// RetrieveRepo returns the repository specified in the configuration. An
// error is returned if one occurs.
func RetrieveRepo(ctx context.Context, cfg *config.Config,
	ghClient *github.Client, redisClient *cache.Codec) (*github.Repository, error) {

	// Check if cached
	var repo *github.Repository

	if err := redisClient.Get(RepoCacheKey, &repo); (err != nil) &&
		(err != cache.ErrCacheMiss) {

		return nil, fmt.Errorf("error retrieving repo from cache: %s",
			err.Error())
	} else if err != cache.ErrCacheMiss {
		// Cached
		return repo, nil
	}

	// Make request
	repo, _, err := ghClient.Repositories.Get(ctx, cfg.GitHub.RepoOwner,
		cfg.GitHub.RepoName)

	if err != nil {
		return nil, fmt.Errorf("error retrieving repository: %s",
			err.Error())
	}

	// Save in cache
	if err := redisClient.Set(&cache.Item{
		Key:        RepoCacheKey,
		Object:     repo,
		Expiration: 0,
	}); err != nil {

		return nil, fmt.Errorf("error saving repository to cache: %s",
			err.Error())
	}

	return repo, nil
}
