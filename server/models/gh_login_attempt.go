package models

import (
	"time"
)

// GitHubLoginAttempt stores information about an GitHub OAuth to prevent cross site request forgery attacks
type GithubLoginAttempt struct {
	// ID is the GitHub login attempt unique database identifier
	ID int

	// CreatedOn stores the date a login attempt was created. Login attempts older than 5 minutes should be deleted
	CreatedOn time.Time

	// State is the unguessable value passed to a GitHub OAuth request to prevent cross site request forgery attacks.
	// This same value is passed to our OAuth callback endpoint and it should match the value passed at the start of the
	// OAuth process. This value should be 32 characters long.
	State string
}
