package github

import (
	"context"
	"fmt"
	"strconv"
)

// Identify returns the user ID of the GitHub user who owns the provided GH auth token
func Identify(ctx context.Context, authToken string) (string, error) {
	// Create GitHub API client
	client := NewUserClient(ctx, authToken)

	// Make identify user request
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return "", fmt.Errorf("error retrieving user information from GitHub API: %s", err.Error())
	}

	return strconv.FormatInt(*(user.ID), 10), nil
}
