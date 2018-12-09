package libissues

import (
	"context"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/auth"
	"github.com/Noah-Huppert/gh-gantt/server/libgh"
	"github.com/Noah-Huppert/gh-gantt/server/libzh"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"github.com/google/go-github/github"
)

// GetCombinedIssuesRequest holds the parameters used to make API requests to GitHub and ZenHub in order to retrieve the
// information for all issues in a repository in Issue form
type GetCombinedIssuesRequest struct {
	// ctx is context
	ctx context.Context

	// logger is used to print debug information
	logger golog.Logger

	// authToken is an API auth token
	authToken auth.AuthToken

	// owner is the name of GitHub user who owns the repository to retrieve issues for
	owner string

	// name is the GitHub repository to retrieve issues for
	name string
}

// NewGetCombinedIssuesRequest creates a GetCombinedIssuesRequest
func NewGetCombinedIssuesRequest(ctx context.Context, logger golog.Logger, authToken auth.AuthToken, owner,
	name string) GetCombinedIssuesRequest {
	return GetCombinedIssuesRequest{
		ctx:       ctx,
		logger:    logger,
		authToken: authToken,
		owner:     owner,
		name:      name,
	}
}

// Do makes the GitHub and ZenHub API requests needed to retrieve information about all issues in a repository.
// Returns a map where keys are issue numbers, and values are issues.
func (r GetCombinedIssuesRequest) Do() (map[int64]*Issue, resp.Responder) {
	// Setup channels which go routines will use to send back various pieces of information
	ghIssuesChan := make(chan github.Issue)
	depsChan := make(chan libzh.ZenHubDependency)
	zhIssuesChan := make(chan libzh.ZenHubBoardIssue)

	// respChan is used to send a responder in the event of an error
	respChan := make(chan resp.Responder)

	// doneChan receives a value from each go routine when it finishes
	doneChan := make(chan bool)

	// Create GitHub client
	client := libgh.NewUserClient(r.ctx, r.authToken.GitHubAuthToken)

	// Get GitHub issues
	go func() {
		// Make API request
		req := libgh.NewListIssuesRequest(r.ctx, client, r.owner, r.name)
		issues, err := req.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(r.logger, http.StatusInternalServerError,
				"error retrieving GitHub issues", err.Error())
			return
		}

		// Send to main thread
		for _, issue := range issues {
			ghIssuesChan <- *issue
		}

		doneChan <- true
	}()

	// Get GitHub repository
	getRepoReq := libgh.NewGetRepositoryRequest(r.ctx, client, r.owner, r.name)

	repo, err := getRepoReq.Do()

	if err != nil {
		return nil, resp.NewStrErrorResponder(r.logger, http.StatusInternalServerError,
			"error retrieving GitHub repository from GitHub API", err.Error())
	}

	// Get ZenHub issue dependencies
	go func() {
		// Make API request
		getDepsReq := libzh.NewGetDependenciesRequest(*(repo.ID), r.authToken.ZenHubAuthToken)

		deps, err := getDepsReq.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(r.logger, http.StatusInternalServerError,
				"error retrieving ZenHub dependencies", err.Error())
			return
		}

		// Send to main thread
		for _, dep := range deps {
			depsChan <- dep
		}

		doneChan <- true
	}()

	// Get ZenHub issue estimates
	go func() {
		// Make API request
		getBoardReq := libzh.NewGetBoardRequest(*(repo.ID), r.authToken.ZenHubAuthToken)

		issues, err := getBoardReq.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(r.logger, http.StatusInternalServerError,
				"error retrieving ZenHub issues", err.Error())
			return
		}

		// Send to main thread
		for _, issue := range issues {
			zhIssuesChan <- issue
		}

		doneChan <- true
	}()

	// Receive results from GitHub and ZenHub API requests
	issues := map[int64]*Issue{}

	numDone := 0

	for numDone < 3 {
		select {
		case resp := <-respChan: // Early endpoint response
			return nil, resp

		case <-doneChan: // GitHub or ZenHub API request go routine done
			numDone++

		case issue := <-ghIssuesChan: // Received issue from GitHub API
			number := int64(*(issue.Number))

			// Check if issue exists in resp
			if _, ok := issues[number]; !ok {
				issues[number] = &Issue{}
			}

			// Save GitHub API issue information in resp
			ci := issues[number]

			ci.Number = int64(number)
			ci.Title = *(issue.Title)
			ci.CreatedAt = *(issue.CreatedAt)

			issues[number] = ci

		case dep := <-depsChan: // Received issue dependency from ZenHub API
			number := int64(dep.Blocked.IssueNumber)

			// Check if issue exists in resp
			if _, ok := issues[number]; !ok {
				issues[number] = &Issue{}
			}

			// Save ZenHub API information in resp
			ci := issues[number]

			ci.Dependencies = append(ci.Dependencies, dep.Blocking.IssueNumber)

			issues[number] = ci

		case issue := <-zhIssuesChan: // Received issue estimate from ZenHub API
			// Check if issue exists in resp
			if _, ok := issues[issue.Number]; !ok {
				issues[issue.Number] = &Issue{}
			}

			// Save ZenHub board information in resp
			ci := issues[issue.Number]

			ci.Estimate = issue.Estimate.Value

			issues[issue.Number] = ci
		}
	}

	// Normalize issues
	for i, issue := range issues {
		// For some reason the ZenHub dependencies API returns dependencies for closed issues, if an issue's title is
		// not set, but it exists in the issues map, then only the ZenHub dependencies API returned it, and the
		// GitHub issues API did not. So we can delete it.
		if len(issue.Title) == 0 {
			delete(issues, i)
			continue
		}

		// If an issue doesn't have an estimate, set it to 0
		if issue.Estimate == 0 {
			ci := issues[i]

			ci.Estimate = 1

			issues[i] = ci
		}

		// Check issue dependencies exist in array
		deps := []int64{}

		for _, dep := range issues[i].Dependencies {
			if _, ok := issues[dep]; ok {
				deps = append(deps, dep)
			}
		}

		ci := issues[i]

		ci.Dependencies = deps

		issues[i] = ci
	}

	return issues, nil
}
