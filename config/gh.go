package config

// GitHub configuration
type GitHub struct {
	// AccessToken is the GitHub API access token used to retrieve issues.
	AccessToken string

	// RepoOwner is the owner of the GitHub repository to retrieve issues from.
	RepoOwner string

	// RepoName is the name of the GitHub repository to retrieve issues from.
	RepoName string
}
