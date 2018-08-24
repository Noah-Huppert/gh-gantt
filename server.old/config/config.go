package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration
type Config struct {
	// Port to listen for HTTP requests on
	Port int `required:"true"`

	// GithubClientID
	GithubClientID string `required:"true" envconfig:"github_client_id"`

	// GithubClientSecret
	GithubClientSecret string `required:"true" envconfig:"github_client_secret"`
}

// NewConfig loads Config values from environment variables. Variables names
// will be capitalized field names from the Config struct, prefixed
// with APP.
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("app", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error loading configuration from "+
			"the environment: %s", err.Error())
	}

	return &cfg, nil
}
