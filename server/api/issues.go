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
	// issuesChan receives GitHub issues from the go routine which calls the GitHub issues API
	issuesChan := make(chan github.Issue)

	// depsChan receives GitHub issue dependency information from the go routine which calls the ZenHub issue
	// dependencies API
	depsChan := make(chan libzh.ZenHubDependency)

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
		issues, err := listIssuesReq.ListIssues()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
				"error retrieving GitHub issues", err.Error())
		}

		for _, issue := range issues {
			issuesChan <- *issue
		}

		doneChan <- true
	}()

	// Get ZenHub issue dependencies
	go func() {
		// Get GitHub repository
		getRepoReq := libgh.NewGetRepositoryRequest(h.ctx, client, request.RepositoryOwner, request.RepositoryName)

		repo, err := getRepoReq.GetRepository()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
				"error retrieving GitHub repository from GitHub API", err.Error())
			return
		}

		// Make ZenHub issue dependencies API request
		getDepsResp := libzh.NewGetDependenciesRequest(*(repo.ID), authToken.ZenHubAuthToken)

		deps, err := getDepsResp.GetDependencies()

		if err != nil {
			respChan <- resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
				"error retrieving ZenHub dependencies", err.Error())
		}

		for _, dep := range deps {
			depsChan <- dep
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

		case issue := <-issuesChan:
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
		}
	}

	// Trim empty issues
	for number, issue := range issuesResp {
		if len(issue.Title) == 0 {
			delete(issuesResp, number)
		}
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"issues": issuesResp,
	}, http.StatusOK)
}
