package zenhub

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"

	"github.com/google/go-github/github"
)

// IssueDeps holds GitHub issue dependency information.
type IssueDeps struct {
	// BlockedBy holds the GitHub issue numbers which are blocking the
	// issue
	BlockedBy []int `json:"blocked_by"`

	// Blocking holds the GitHub issue numbers which are being blocked
	// by the issue
	Blocking []int `json:"blocking"`
}

// DepIssue is a github.Issue with fields for dependency information
type DepIssue struct {
	github.Issue
	IssueDeps
}

// NewDepIssue creates a new DepIssue instance from a github.Issue and a IssueDeps
func NewDepIssue(iss github.Issue, deps IssueDeps) DepIssue {
	return DepIssue{
		iss,
		deps,
	}
}

// DepsURL is the URL used to retrieve issue dependency information.
// Expected to be used in a fmt.*f function of some kind. Has 2 template
// values, both numbers:
// 	1. repo id
//	2. issue id
const DepsURL string = "https://api.zenhub.io/v4/repositories/%d/issues/%d/dependencies"

// extractIssueNumbers returns an array of issue numbers from the specified
// array of maps. An error is returned if one occurs.
func extractIssueNumbers(data []map[string]interface{}) ([]int, error) {
	numbers := []int{}

	// Loop through items
	for _, item := range data {
		// Convert to int
		val := item["issue_number"]
		num, ok := val.(float64)
		if !ok {
			return nil, fmt.Errorf("error converting value to string, "+
				"val: %#v", val)
		}

		numbers = append(numbers, int(num))
	}

	return numbers, nil
}

// RetrieveDeps returns an IssueDeps instance containing dependency information
// for the specified issue. An error is returned if one occurs.
func RetrieveDeps(ctx context.Context, cfg *config.Config, repoId int64,
	issueId int) (IssueDeps, error) {

	// Setup ZenHub API request
	reqUrl := fmt.Sprintf(DepsURL, repoId, issueId)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return IssueDeps{}, fmt.Errorf("error creating HTTP request: "+
			"%s", err.Error())
	}

	req.Header.Add("x-authentication-token", cfg.ZenHub.APIToken)

	// Make request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return IssueDeps{}, fmt.Errorf("error making HTTP request: %s",
			err.Error())
	}
	defer resp.Body.Close()

	// Read request body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return IssueDeps{}, fmt.Errorf("error reading HTTP response "+
			"body: %s", err.Error())
	}

	body := map[string][]map[string]interface{}{}
	if err = json.Unmarshal(bodyBytes, &body); err != nil {
		return IssueDeps{}, fmt.Errorf("error unmarshalling HTTP "+
			"response into JSON: %s", err.Error())
	}

	// Check request succeeded
	if resp.StatusCode == http.StatusOK {
		deps := IssueDeps{}

		// Blocked by
		blockedBy, err := extractIssueNumbers(body["blocked_by"])
		if err != nil {
			return IssueDeps{}, fmt.Errorf("error retrieving "+
				"blocked by issue numbers: %s", err.Error())
		}

		deps.BlockedBy = blockedBy

		// Blocking
		blocking, err := extractIssueNumbers(body["blocking"])
		if err != nil {
			return IssueDeps{}, fmt.Errorf("error retrieving "+
				"blocking issue numbers: %s", err.Error())
		}

		deps.Blocking = blocking

		return deps, nil
	} else {
		return IssueDeps{}, fmt.Errorf("non OK status code response: "+
			"%d, body: %#v", resp.StatusCode, body)
	}
}
