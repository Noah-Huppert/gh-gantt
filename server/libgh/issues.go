package libgh

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// ListIssuesRequest holds the parameters for a list issues GitHub API request
type ListIssuesRequest struct {
	// ctx is context
	ctx context.Context

	// client is a GitHub API client
	client *github.Client

	// owner is the name of the GitHub user who owns the repository to list issues for
	owner string

	// name of the GitHub repository to list issues for
	name string
}

// NewListIssuesRequest creates a new ListIssuesRequest
func NewListIssuesRequest(ctx context.Context, client *github.Client, owner, name string) ListIssuesRequest {
	return ListIssuesRequest{
		ctx:    ctx,
		client: client,
		owner:  owner,
		name:   name,
	}
}

// ListIssues makes the list issues GitHub API request
func (r ListIssuesRequest) Do() ([]*github.Issue, error) {
	issues, _, err := r.client.Issues.ListByRepo(r.ctx, r.owner, r.name, nil)

	if err != nil {
		return nil, fmt.Errorf("error retrieving GitHub issues: %s", err.Error())
	}

	return issues, nil
}
