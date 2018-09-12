package models

import (
	"fmt"
	"time"

	"github.com/dchest/uniuri"
	"github.com/jmoiron/sqlx"
)

// GitHubLoginAttempt stores information about an GitHub OAuth to prevent cross site request forgery attacks
type GitHubLoginAttempt struct {
	// ID is the GitHub login attempt unique database identifier
	ID int

	// CreatedOn stores the date a login attempt was created. Login attempts older than 5 minutes should be deleted
	CreatedOn time.Time

	// State is the unguessable value passed to a GitHub OAuth request to prevent cross site request forgery attacks.
	// This same value is passed to our OAuth callback endpoint and it should match the value passed at the start of the
	// OAuth process. This value should be 32 characters long.
	State string
}

// NewGitHubLoginAttempt creates a new GitHubLoginAttempt.
// The CreatedOn time set to the current time and the State field set to a random 32 character long string. The ID
// field will remain unset.
func NewGitHubLoginAttempt() GitHubLoginAttempt {
	return GitHubLoginAttempt{
		CreatedOn: time.Now(),
		State:     uniuri.NewLen(32),
	}
}

// QueryByState attempts to find a GitHubLoginAttempt in the database with a matching State field. Returns sql.ErrNoRows
// if a matching row is not found.
func (a *GitHubLoginAttempt) QueryByState(db *sqlx.DB) error {
	return db.Get(a, "SELECT id, created_on FROM github_login_attempts WHERE state = ?", a.State)
}

// Insert login attempt into the database
func (a GitHubLoginAttempt) Insert(db *sqlx.DB) error {
	// Start transaction
	tx, err := db.Begin()

	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}

	// Insert
	_, err = tx.Exec("INSERT INTO github_login_attempts (created_on, state) VALUES(?, ?)", a.CreatedOn, a.State)

	if err != nil {
		return fmt.Errorf("error inserting GitHub login attempt: %s", err.Error())
	}

	// Commit transaction
	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("error committing transaction: %s", err.Error())
	}

	// Done
	return nil
}
