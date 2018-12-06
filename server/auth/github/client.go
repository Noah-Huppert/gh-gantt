package github

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewUserClient creates a new GitHub API client with a user's GitHub authentication token
func NewUserClient(ctx context.Context, githubAuthToken string) *github.Client {
	tokenClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: githubAuthToken,
	}))

	return github.NewClient(tokenClient)
}
