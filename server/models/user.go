package models

// User represents a GitHub user who has logged in to use the GH Gantt application
type User struct {
	// GitHubID is user's ID which was retrieved from the GitHub API
	GitHubID string

	// Name is what the user refers to themselves as
	Name string

	// Login is the user's GitHub username
	Login string

	// ProfilePictureURL is the URL of the user's GitHub profile picture
	ProfilePictureURL string
}
