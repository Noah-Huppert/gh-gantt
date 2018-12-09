package libzh

import (
	"fmt"
)

// GetBoardRequest holds parameters for a get board ZenHub API request
type GetBoardRequest struct {
	// repositoryID is the ID of the repository to fetch a ZenHub board for
	repositoryID int64

	// authToken is a ZenHub API auth token
	authToken string
}

// NewGetBoardRequest creates a GetBoardRequest
func NewGetBoardRequest(repositoryID int64, authToken string) GetBoardRequest {
	return GetBoardRequest{
		repositoryID: repositoryID,
		authToken:    authToken,
	}
}

// ZenHubBoardIssue holds information about a GitHub issue on a ZenHub board
type ZenHubBoardIssue struct {
	// Number is the ID of an issue, unique inside of a repository
	Number int64 `json:"issue_number"`

	// Estimate is a value used to indicate how much work an issue will take to resolve
	Estimate struct {
		// Value is the estimate value
		Value int64 `json:"value"`
	} `json:"estimate"`
}

// ZenHubBoardResponse is the value returned by a get board ZenHub API request
type ZenHubBoardResponse struct {
	// Pipelines holds a list of ZenHub pipelines on the board
	Pipelines []struct {
		// Issues holds a list of issues in the pipeline
		Issues []ZenHubBoardIssue `json:"issues"`
	} `json:"pipelines"`
}

// Do makes a get board ZenHub API request
func (r GetBoardRequest) Do() ([]ZenHubBoardIssue, error) {
	// Make API request
	req := ZenHubAPIRequest{
		url:       fmt.Sprintf("/p1/repositories/%d/board", r.repositoryID),
		authToken: r.authToken,
	}

	var resp ZenHubBoardResponse

	err := req.Do(&resp)

	if err != nil {
		return nil, fmt.Errorf("error making get board ZenHub API request: %s", err.Error())
	}

	// Aggregate issues
	issues := []ZenHubBoardIssue{}

	for _, pipeline := range resp.Pipelines {
		for _, issue := range pipeline.Issues {
			issues = append(issues, issue)
		}
	}

	return issues, nil
}
