package config

// AuthConfig holds user authentication configuration
type AuthConfig struct {
	// ServiceName is the name of the GH Gantt server used in the AuthToken Issuer and Audience field
	ServiceName string `required:"true" envconfig:"service_name"`

	// SigningSecret is used to sign JWTs
	SigningSecret string `required:"true" envconfig:"signing_secret"`

	// GHStateSigningPrivKey is the ed25519 key used to sign the state field in a GH authentication request.
	// Not in the normal OpenSSH key format. Instead the first 32 bits are the private key and the last 32 bits are the
	// public key.
	GHStateSigningKey string `validate:"len=64" required:"true" envconfig:"gh_state_signing_key"`
}

// GetGHStateSigningPubKey extracts the public key from the GHStateSigningKey field
func (c AuthConfig) GetGHStateSigningPubKey() []byte {
	return []byte(c.GHStateSigningKey)[32:]
}

// GetGHStateSigningPrivKey extracts the private key from the GHStateSigningKey field. In the Go ed25519 library the
// private key and public key are considered the "private key". So this method simply coverts the entire field to bytes.
func (c AuthConfig) GetGHStateSigningPrivKey() []byte {
	return []byte(c.GHStateSigningKey)
}
