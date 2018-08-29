package config

// HTTPConfig holds HTTP related application configuration
type HTTPConfig struct {
	// Port to listen for HTTP requests on
	// Set by APP_PORT environment variable
	Port int `required:"true"`
}
