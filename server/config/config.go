package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config holds application configuration
type Config struct {
	// Port to listen for HTTP requests on
	Port int `required:"true"`

	// GithubClientID is the server's GitHub application's client ID
	GithubClientID string `required:"true" envconfig:"github_client_id"`

	// GithubClientSecret is the server's GitHub application's client secret
	GithubClientSecret string `required:"true" envconfig:"github_client_secret"`

	// DBHost is the host of the database
	DBHost string `required:"true" envconfig:"db_host"`

	// DBName is the name of the database to save data in
	DBName string `required:"true" envconfig:"db_name"`

	// DBUsername is the username used to authenticate with the database
	DBUsername string `required:"true" envconfig:"db_username"`

	// DBPassword is ithe password used to authenticate with the database
	DBPassword string `envconfig:"db_password"`
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
