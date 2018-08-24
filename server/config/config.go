package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration
type Config struct {
	// Port to listen for HTTP requests on
	Port int `required:"true"`
}

// NewConfig loads Config values from environment variables. Variables names
// will be capitalized field names from the Config struct, prefixed
// with GH_GANTT.
func NewConfig() (*Config, error) {
	var cfg Config

	err := envconfig.Process("gh_gantt", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error loading configuration from "+
			"the environment: %s", err.Error())
	}

	return &cfg, nil
}
