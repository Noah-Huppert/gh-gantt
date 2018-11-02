package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/validator.v2"
)

// Config holds application configuration
type Config struct {
	// HTTPConfig holds HTTP related application configuration
	HTTPConfig

	// GitHubConfig holds GitHub related application configuration
	GitHubConfig

	// AuthConfig holds user authentication related application configuration
	AuthConfig
}

// NewConfig loads Config values from environment variables. Variables names
// will be capitalized field names from the Config struct, prefixed
// with APP.
func NewConfig() (*Config, error) {
	var cfg Config

	// Load from environment
	err := envconfig.Process("app", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error loading configuration from the environment: %s", err.Error())
	}

	// Validate
	err = validator.Validate(cfg)
	if err != nil {
		return nil, fmt.Errorf("error validating configuration: %s", err.Error())
	}

	return &cfg, nil
}
