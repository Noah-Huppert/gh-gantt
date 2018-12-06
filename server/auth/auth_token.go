package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// AuthToken is issued to a user to prove they are authenticated
type AuthToken struct {
	// Issuer is the name of the service which issued the token
	Issuer string

	// Audience is the name of the service who is meant to received the token
	Audience string

	// GitHubUserID is the ID of the authenticated GitHub user, used as the subject field in the JWT
	GitHubUserID string

	// GitHubAuthToken is the user's GitHub authentication token
	GitHubAuthToken string
}

// NewAuthToken creates a new AuthToken.
// The serviceName is used for the Issuer and Audience field
func NewAuthToken(serviceName, ghUserID, ghAuthToken string) AuthToken {
	return AuthToken{
		Issuer:          serviceName,
		Audience:        serviceName,
		GitHubUserID:    ghUserID,
		GitHubAuthToken: ghAuthToken,
	}
}

// claims returns a map of JWT claims to encode
func (t AuthToken) claims() map[string]interface{} {
	return map[string]interface{}{
		"iss":               t.Issuer,
		"aud":               t.Audience,
		"sub":               t.GitHubUserID,
		"github_auth_token": t.GitHubAuthToken,
	}
}

// Encode an auth token into a string
func (t AuthToken) Encode(signingSecret string) (string, error) {
	// Generate claims
	var claims jwt.MapClaims = t.claims()

	// Encode
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(signingSecret))
	if err != nil {
		return "", fmt.Errorf("error encoding JWT: %s", err.Error())
	}

	return tokenStr, nil
}
