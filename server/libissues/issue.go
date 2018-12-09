package libissues

import (
	"time"
)

// Issue holds information about a GitHub issue from the GitHub API and the ZenHub API
type Issue struct {
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
