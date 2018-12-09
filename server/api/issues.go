package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/libgh"
	"github.com/Noah-Huppert/gh-gantt/server/libzh"
	"github.com/Noah-Huppert/gh-gantt/server/req"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"github.com/google/go-github/github"
)

// IssuesHandler implements resp.ResponderHandler by returning a list of GitHub issues for a repository
type IssuesHandler struct {
	// ctx is application context
	ctx context.Context

	// logger prints debug information
	logger golog.Logger

	// cfg is application configuration
	cfg config.Config
}

// IssuesRequest is the format of a issues request
type IssuesRequest struct {
	// RepositoryOwner is the repository's owner
	RepositoryOwner string `json:"repository_owner" validate:"nonzero"`

	// RepositoryName is the repository's name
	RepositoryName string `json:"repository_name" validate:"nonzero"`
}

// combinedIssue holds information about a GitHub issue from the GitHub API and the ZenHub API
type combinedIssue struct {
	// Number is an ID used to identify an issue, unique only within it's GitHub repository
	Number int64 `json:"number"`

	// Title is the issue's title
	Title string `json:"title"`

	// CreatedAt is the date and time the issue was created
	CreatedAt time.Time `json:"created_at"`

	// Estimate value for issue
	Estimate int64 `json:"estimate"`

	// Dependencies holds a list of issue numbers which the issue is blocked by
	Dependencies []int64 `json:"dependencies"`
}

// NewIssuesHandler creates a new IssuesHandler
func NewIssuesHandler(ctx context.Context, logger golog.Logger, cfg config.Config) IssuesHandler {
	return IssuesHandler{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h IssuesHandler) Handle(r *http.Request) resp.Responder {
	// Check auth token
	authToken, errResp := req.CheckAuthToken(h.logger, h.cfg, r)
	if errResp != nil {
		return errResp
	}

	// Get query parameters
	var request IssuesRequest

	repoOwner := r.URL.Query().Get("repository_owner")
	if len(repoOwner) == 0 {
		return resp.NewStrErrorResponder(h.logger, http.StatusBadRequest,
			"repository_owner query parameter must be provided", "")
	}

	repoName := r.URL.Query().Get("repository_name")
	if len(repoOwner) == 0 {
		return resp.NewStrErrorResponder(h.logger, http.StatusBadRequest,
			"repository_name query parameter must be provided", "")
	}

	request.RepositoryOwner = repoOwner
	request.RepositoryName = repoName

	// Setup channels
	// ghIssuesChan receives GitHub issues from the go routine which calls the GitHub issues API
	ghIssuesChan := make(chan github.Issue)

	// depsChan receives GitHub issue dependency information from the go routine which calls the ZenHub issue
	// dependencies API
	depsChan := make(chan libzh.ZenHubDependency)

	// zhIssuesChan receives ZenHub issues from the go routine which calls the ZenHub board API
	zhIssuesChan := make(chan libzh.ZenHubBoardIssue)

	// respChan receives responders from either of the 2 go routines mentioned above. If a responder is received then
	// it is immediately returned to the client.
	respChan := make(chan resp.Responder)

	// doneChan receives data from both of the 2 go routines mentioned above when they are finished sending their data
	// back to the main thread.
	doneChan := make(chan bool)

	// Create GitHub client
	client := libgh.NewUserClient(h.ctx, authToken.GitHubAuthToken)

	// Get GitHub issues
	go func() {
		listIssuesReq := libgh.NewListIssuesRequest(h.ctx, client, request.RepositoryOwner, request.RepositoryName)
		issues, err := listIssuesReq.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
				"error retrieving GitHub issues", err.Error())
			return
		}

		for _, issue := range issues {
			ghIssuesChan <- *issue
		}

		doneChan <- true
	}()

	// Get GitHub repository
	getRepoReq := libgh.NewGetRepositoryRequest(h.ctx, client, request.RepositoryOwner, request.RepositoryName)

	repo, err := getRepoReq.Do()

	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error retrieving GitHub repository from GitHub API", err.Error())
	}

	// Get ZenHub issue dependencies
	go func() {
		// Make ZenHub issue dependencies API request
		getDepsReq := libzh.NewGetDependenciesRequest(*(repo.ID), authToken.ZenHubAuthToken)

		deps, err := getDepsReq.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
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
		// Make ZenHub get board API request
		getBoardReq := libzh.NewGetBoardRequest(*(repo.ID), authToken.ZenHubAuthToken)

		issues, err := getBoardReq.Do()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
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
	issuesResp := map[int64]combinedIssue{}

	numDone := 0

	for numDone < 2 {
		select {
		case resp := <-respChan:
			return resp

		case <-doneChan:
			numDone++

		case issue := <-ghIssuesChan:
			number := int64(*(issue.Number))

			// Check if issue exists in resp
			if _, ok := issuesResp[number]; !ok {
				issuesResp[number] = combinedIssue{}
			}

			// Save GitHub API issue information in resp
			ci := issuesResp[number]

			ci.Number = int64(number)
			ci.Title = *(issue.Title)
			ci.CreatedAt = *(issue.CreatedAt)

			issuesResp[number] = ci

		case dep := <-depsChan:
			number := int64(dep.Blocked.IssueNumber)

			// Check if issue exists in resp
			if _, ok := issuesResp[number]; !ok {
				issuesResp[number] = combinedIssue{}
			}

			// Save ZenHub API information in resp
			ci := issuesResp[number]

			ci.Dependencies = append(ci.Dependencies, dep.Blocking.IssueNumber)

			issuesResp[number] = ci

		case issue := <-zhIssuesChan:
			// Check if issue exists in resp
			if _, ok := issuesResp[issue.Number]; !ok {
				issuesResp[issue.Number] = combinedIssue{}
			}

			// Save ZenHub board information in resp
			ci := issuesResp[issue.Number]

			ci.Estimate = issue.Estimate.Value

			issuesResp[issue.Number] = ci
		}
	}

	// Normalize issues
	for i, issue := range issuesResp {
		// For some reason the ZenHub dependencies API returns dependencies for closed issues, if an issue's title is
		// not set, but it exists in the issues map, then only the ZenHub dependencies API returned it, and the
		// GitHub issues API did not. So we can delete it.
		if len(issue.Title) == 0 {
			delete(issuesResp, i)
			continue
		}

		// If an issue doesn't have an estimate, set it to 0
		if issue.Estimate == 0 {
			ci := issuesResp[i]

			ci.Estimate = 1

			issuesResp[i] = ci
		}
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"issues": issuesResp,
	}, http.StatusOK)
}
