package libgh

import (
	"context"
	"fmt"
	"strconv"
)

// IdentifyAuthTokenRequest holds the parameters of a GitHub API request to identify the owner of a GitHub
// authentication token
type IdentifyAuthTokenRequest struct {
	// ctx is context
	ctx context.Context

	// authToken is the authentication token to identify
	authToken string
}

// NewIdentifyAuthTokenRequest creates a new IdentifyAuthTokenRequest
func NewIdentifyAuthTokenRequest(ctx context.Context, authToken string) IdentifyAuthTokenRequest {
	return IdentifyAuthTokenRequest{
		ctx:       ctx,
		authToken: authToken,
	}
}

// Identify returns the user ID of the GitHub user who owns the provided GH auth token
func (r IdentifyAuthTokenRequest) Identify() (string, error) {
	// Create GitHub API client
	client := NewUserClient(r.ctx, r.authToken)

	// Make identify user request
	user, _, err := client.Users.Get(r.ctx, "")
	if err != nil {
		return "", fmt.Errorf("error retrieving user information from GitHub API: %s", err.Error())
	}

	return strconv.FormatInt(*(user.ID), 10), nil
}
