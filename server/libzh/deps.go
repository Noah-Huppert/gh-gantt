package libzh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (r GetDependenciesRequest) GetDependencies() ([]ZenHubDependency, error) {
	// Setup request
	depsReqURL, err := url.Parse(fmt.Sprintf("https://api.zenhub.io/p1/repositories/%d/dependencies", r.RepositoryID))

	if err != nil {
		return nil, fmt.Errorf("error parsing request URL: %s", err.Error())
	}

	depsReq := &http.Request{
		Method: http.MethodGet,
		URL:    depsReqURL,
		Header: map[string][]string{
			"X-Authentication-Token": []string{r.ZenHubAuthToken},
		},
	}

	// Make request
	depsResp, err := http.DefaultClient.Do(depsReq)

	if err != nil {
		return nil, fmt.Errorf("error making get dependencies ZenHub API request: %s", err.Error())
	}

	if depsResp.StatusCode != http.StatusOK {
		// If response error, read body to print
		respBody, err := ioutil.ReadAll(depsResp.Body)

		if err != nil {
			return nil, fmt.Errorf("get dependencies ZenHub API response status not OK, and an error occurred reading "+
				"the body: %s", err.Error())
		}

		return nil, fmt.Errorf("get dependencies ZenHub API response status not OK, status: %s, body: %s", depsResp.Status, string(respBody))
	}

	// Decode ZenHub issue dependencies API response
	var deps ZenHubDependenciesResponse

	decoder := json.NewDecoder(depsResp.Body)

	err = decoder.Decode(&deps)
	if err != nil {
		return nil, fmt.Errorf("error decoding get dependencies ZenHub API response body: %s", err.Error())
	}

	return deps.Dependencies, nil
}
