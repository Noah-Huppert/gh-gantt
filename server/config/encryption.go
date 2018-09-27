package config

// EncryptionConfig holds symmetric encryption related configuration
type EncryptionConfig struct {
	// AuthTokenEncryptionSecret is the secret key used to symmetrically encrypt third party authentication tokens
	AuthTokenEncryptionSecret string `required:"true" envconfig:"auth_token_encryption_secret"`
}
