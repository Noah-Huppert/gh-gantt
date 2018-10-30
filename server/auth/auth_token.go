package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
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

// claims returns a map of JWT claims to encode
func (t AuthToken) claims() map[string]string {
	return map[string]string{
		"iss":               t.Issuer,
		"aud":               t.Audience,
		"sub":               t.GitHubUserID,
		"github_auth_token": t.GitHubAuthToken,
	}
}

func (t AuthToken) Encode(signingSecret string) (string, error) {
	// Generate claims
	var claims jwt.MapClaims = structs.Map(t)

	// Encode
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(signingSecret))
	if err != nil {
		return "", fmt.Errorf("error encoding JWT: %s", err.Error())
	}

	return tokenStr, nil
}
