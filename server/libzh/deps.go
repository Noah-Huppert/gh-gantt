package libzh

import (
	"fmt"
)

// GetDependenciesRequest hold the parameters for a get ZenHub dependencies API request
type GetDependenciesRequest struct {
	// RepositoryID is the ID of the repository to get issue dependencies for
	RepositoryID int64

	// ZenHubAuthToken is the ZenHub API token
	ZenHubAuthToken string
}

// NewGetDependenciesRequest creates a new GetDependenciesRequest
func NewGetDependenciesRequest(repositoryID int64, zenhubAuthToken string) GetDependenciesRequest {
	return GetDependenciesRequest{
		RepositoryID:    repositoryID,
		ZenHubAuthToken: zenhubAuthToken,
	}
}

// ZenHubIssueID identifies an issue
type ZenHubIssueID struct {
	// RepositoryID is the ID of the repository the issue belongs to
	RepositoryID int64 `json:"repo_id"`

	// IssueNumber is the ID of the issue
	IssueNumber int64 `json:"issue_number"`
}

// ZenHubDependency holds information about an issue dependency from ZenHub
type ZenHubDependency struct {
	// Blocking holds information about the issue which is doing the blocking
	Blocking ZenHubIssueID `json:"blocking"`

	// Blocked holds information about the issue being blocked
	Blocked ZenHubIssueID `json:"blocked"`
}

// ZenHubDependenciesResponse holds a list of issue dependencies from ZenHub
type ZenHubDependenciesResponse struct {
	// Dependencies is a list of issue dependencies from ZenHub
	Dependencies []ZenHubDependency `json:"dependencies"`
}

// Do makes the get dependencies ZenHub API request
func (r GetDependenciesRequest) Do() ([]ZenHubDependency, error) {
	req := ZenHubAPIRequest{
		url:       fmt.Sprintf("https://api.zenhub.io/p1/repositories/%d/dependencies", r.RepositoryID),
		authToken: r.ZenHubAuthToken,
	}

	var depsResp ZenHubDependenciesResponse

	err := req.Do(&depsResp)

	if err != nil {
		return nil, fmt.Errorf("error making get dependencies ZenHub API request: %s", err.Error())
	}

	return depsResp.Dependencies, nil
}
