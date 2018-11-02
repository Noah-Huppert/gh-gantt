package config

// GitHubConfig holds GitHub related application configuration
type GitHubConfig struct {
	// GitHubClientID is the server's GitHub application's client ID
	// Set via the APP_GITHUB_CLIENT_ID environment variable
	GitHubClientID string `required:"true" envconfig:"github_client_id"`

	// GitHubClientSecret is the server's GitHub application's client secret
	// Set via the APP_GITHUB_CLIENT_SECRET environment variable
	GitHubClientSecret string `required:"true" envconfig:"github_client_secret"`
}
