package zenhub

import (
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
