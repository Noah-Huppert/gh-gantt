package config

// AuthConfig holds user authentication configuration
type AuthConfig struct {
	// SigningSecret is used to sign JWTs
	SigningSecret string `required:"true" envconfig:"signing_secret"`

	// GHStateSigningPrivKey is the ed25519 private key used to sign the state field in a GH authentication request
	GHStateSigningPrivKey []byte `required:"true" envconfig:"gh_state_signing_priv_key"`

	// GHStateSigningPubKey is the ed25519 public key used to sign the state field in a GH authentication request
	GHStateSigningPubKey []byte `required:"true" envconfig:"gh_state_signing_pub_key"`
}
