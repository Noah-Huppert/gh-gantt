package config

// AuthConfig holds user authentication configuration
type AuthConfig struct {
	// SigningSecret is used to sign JWTs
	SigningSecret string `required:"true" envconfig:"signing_secret"`
}
