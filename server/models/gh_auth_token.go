package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// GitHubAccessToken is a GitHub API access token
type GitHubAccessToken struct {
	// ID is a unique identifier
	ID int64

	// GitHubUserID is the GitHub ID of the user who the access token is for
	GitHubUserID string

	// EncryptedAccessToken is the encrypted GitHub access token value
	EncryptedAccessToken string
}

// Insert an encrypted GitHubAccessToken. Saves the inserted row's ID in the ID field
func (t *GitHubAccessToken) Insert(db sqlx.DB) error {
	// Begin transaction
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("error starting transaction: %s", err.Error())
	}

	// Insert
	err := tx.QueryRowx("INSERT INTO github_authentication_tokens (github_user_id, encrypted_access_token) VALUES "+
		"($1, $2) RETURNING id", t.GitHubUserID, t.EncryptedAccessToken).StructScan(t.ID)
	if err != nil {
		return fmt.Errorf("error executing insert query: %s", err.Error())
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %s", err.Error())
	}

	return nil
}
