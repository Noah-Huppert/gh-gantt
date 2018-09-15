package models

import (
	"github.com/jmoiron/sqlx"
)

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

// NewUser creates a new User
func NewUser(ghID, name, login, profilePicURL string) *User {
	return &User{
		GitHubID:          ghID,
		Name:              name,
		Login:             login,
		ProfilePictureURL: profilePicURL,
	}
}

// QueryByGitHubID searches the database for a User with a matching GitHubID.
//
// Returns an error if one occurs. Returns sql.ErrNoRows if a matching User is not found.
func (u *User) QueryByGitHubID(db *sqlx.DB) error {
	return db.Get(u, "SELECT name, login, profile_picture_url FROM users WHERE github_id = ?", u.GitHubID)
}
