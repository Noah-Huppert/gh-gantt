package models

// GitHubAccessToken is a GitHub API access token
type GitHubAccessToken struct {
	// ID is a unique identifier
	ID int64

	// GitHubUserID is the GitHub ID of the user who the access token is for
	GitHubUserID string

	// EncryptedAccessToken is the encrypted GitHub access token value
	EncryptedAccessToken string
}
