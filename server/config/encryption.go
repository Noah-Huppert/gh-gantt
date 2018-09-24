package config

// EncryptionConfig holds symmetric encryption related configuration
type EncryptionConfig struct {
	// ThirdPartyAuthTokenEncryptionSecret is the secret key used to symmetrically encrypt third party
	// authentication tokens
	ThirdPartyAuthTokenEncryptionSecret string `required:"true" envconfig:"third_party_auth_token_encryption_secret"`
}
