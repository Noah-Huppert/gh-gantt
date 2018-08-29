package config

// GitHubConfig holds GitHub related application configuration
type GitHubConfig struct {
	// GithubClientID is the server's GitHub application's client ID
	// Set via the APP_GITHUB_CLIENT_ID environment variable
	GithubClientID string `required:"true" envconfig:"github_client_id"`

	// GithubClientSecret is the server's GitHub application's client secret
	// Set via the APP_GITHUB_CLIENT_SECRET environment variable
	GithubClientSecret string `required:"true" envconfig:"github_client_secret"`
}
