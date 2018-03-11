package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// instance caches the application configuration object so it only needs to be
// loaded once.
var instance *Config

// Load returns the application configuration file contents.
func Load() (*Config, error) {
	// Check if loaded
	if instance != nil {
		return instance, nil
	}

	// Load
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading configuration: %s",
			err.Error())
	}

	instance = &Config{}
	if err := viper.Unmarshal(instance); err != nil {
		return nil, fmt.Errorf("error unmarshalling configuration: %s",
			err.Error())
	}

	return instance, nil
}
