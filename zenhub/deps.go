package zenhub

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"

	"github.com/Noah-Huppert/gh-gantt/config"
)

// IssueDeps holds GitHub issue dependency information.
type IssueDeps struct {
	// ID is the GitHub issue identifier
	ID int64

	// BlockedBy holds the IDs of GitHub issues which are blocking the
	// issue specified by `ID`.
	BlockedBy []int64

	// Blocking holds the IDs of the GitHub issues which are being blocked
	// by the issue specified by `ID`.
	Blocking []int64
}

// DepsURL is the URL used to retrieve issue dependency information.
// Expected to be used in a fmt.*f function of some kind. Has 2 template
// values, both numbers:
// 	1. repo id
//	2. issue id
const DepsURL string = "api.zenhub.io/v4/repositories/%d/issues/%d/dependencies"

// RetrieveDeps returns an IssueDeps instance containing dependency information
// for the specified issue. An error is returned if one occurs.
func RetrieveDeps(ctx context.Context, cfg *config.Config,
	ghClient *github.Client, repoId int64, issueId int64) (IssueDeps, error) {

	// Setup ZenHub API request
	reqUrl := fmt.Sprintf(DepsURL, repoId, issueId)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return IssueDeps{}, fmt.Errorf("error creating HTTP request: "+
			"%s", err.Error())
	}

	req.Header.Set("X-Authentication-Token", cfg.ZenHub.APIToken)

	// Make request
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return IssueDeps{}, fmt.Errorf("error making HTTP request: %s",
			err.Error())
	}
	defer resp.Body.Close()

	// Check request succeeded
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("body: %#v", resp.Body)
		return IssueDeps{}, nil
	} else {
		return IssueDeps{}, fmt.Errorf("non OK status code response: "+
			"%d, body: %#v", resp.StatusCode, resp.Body)
	}
}
