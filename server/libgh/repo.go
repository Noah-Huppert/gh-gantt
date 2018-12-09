package libgh

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// GetRepositoryRequest are the parameters for a get repository GitHub API request
type GetRepositoryRequest struct {
	// ctx is context
	ctx context.Context

	// client is a GitHub API client
	client *github.Client

	// owner is the name of the GitHub user who owns the repository
	owner string

	// name is the GitHub repository to retrieve information about
	name string
}

// NewGetRepositoryRequest creates a new GetRepositoryRequest
func NewGetRepositoryRequest(ctx context.Context, client *github.Client, owner, name string) GetRepositoryRequest {
	return GetRepositoryRequest{
		ctx:    ctx,
		client: client,
		owner:  owner,
		name:   name,
	}
}

// GetRepository makes a get repository GitHub API request
func (r GetRepositoryRequest) Do() (*github.Repository, error) {
	repo, _, err := r.client.Repositories.Get(r.ctx, r.owner, r.name)

	if err != nil {
		return nil, fmt.Errorf("error making get repository GitHub API request: %s", err.Error())
	}

	return repo, nil
}
