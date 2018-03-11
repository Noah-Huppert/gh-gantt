package gh

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"

	"github.com/Noah-Huppert/gh-gantt/config"
)

// NewClient creates a new GitHub API client. An error is returned if one occurs.
func NewClient(ctx context.Context, cfg *config.Config) *github.Client {
	// OAuth
	tokenSrc := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.GitHub.AccessToken},
	)
	tokenClient := oauth2.NewClient(ctx, tokenSrc)

	// GH
	client := github.NewClient(tokenClient)
	return client
}
