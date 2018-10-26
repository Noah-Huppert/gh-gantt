package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration
type Config struct {
	// HTTPConfig holds HTTP related application configuration
	HTTPConfig

	// GitHubConfig holds GitHub related application configuration
	GitHubConfig

	// DBConfig holds database related application configuration
	DBConfig

	// AuthConfig holds user authentication related application configuration
	AuthConfig
}

// NewConfig loads Config values from environment variables. Variables names
// will be capitalized field names from the Config struct, prefixed
// with APP.
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("app", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error loading configuration from the environment: %s", err.Error())
	}

	return &cfg, nil
}
