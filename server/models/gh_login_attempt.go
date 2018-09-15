package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/dchest/uniuri"
	"github.com/jmoiron/sqlx"
)

// GitHubLoginAttempt stores information about an GitHub OAuth to prevent cross site request forgery attacks
type GitHubLoginAttempt struct {
	// ID is the GitHub login attempt unique database identifier
	ID int64

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
func NewGitHubLoginAttempt() &GitHubLoginAttempt {
	return &GitHubLoginAttempt{
		CreatedOn: time.Now(),
		State:     uniuri.NewLen(32),
	}
}

// QueryByState attempts to find a GitHubLoginAttempt in the database with a matching State field. Returns sql.ErrNoRows
// if a matching row is not found.
func (a *GitHubLoginAttempt) QueryByState(db *sqlx.DB) error {
	return db.Get(a, "SELECT id, created_on FROM github_login_attempts WHERE state = ?", a.State)
}

// Insert a GitHub login attempt into the database. The new row's id will be saved under the ID field
func (a *GitHubLoginAttempt) Insert(db *sqlx.DB) error {
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}

	// Insert
	res, err = tx.Exec("INSERT INTO github_login_attempts (created_on, state) VALUES(?, ?)", a.CreatedOn, a.State)
	if err != nil {
		return fmt.Errorf("error executing insert statement: %s", err.Error())
	}

	// Record inserted row's id in struct
	id, err := a.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving the inserted row's id: %s", err.Error())
	}

	a.ID = id

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %s", err.Error())
	}

	// Done
	return nil
}

// Delete login attempt from database based on the ID field.
//
// An error is returned if the GitHub login attempt could not be deleted. sql.ErrNoRows is returned if no GitHub login
// attempt with the specified ID exists.
func (a GitHubLoginAttempt) Delete(db *sqlx.DB) error {
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}

	// Delete
	res, err := tx.Exec("DELETE FROM github_login_attempts WHERE id = ?", a.ID)
	if err != nil {
		return fmt.Errorf("error executing delete statement: %s", err.Error())
	}

	// Check that a row was deleted
	numDel, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("error determining how many rows were deleted: %s", err.Error())
	}

	if numDel == 0 {
		return sql.ErrNoRows
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %s", err.Error())
	}

	// Done
	return nil
}
