package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/fatih/structs"
)

// AuthToken is issued to a user to prove they are authenticated
type AuthToken struct {
	// GitHubUserID is the ID of the authenticated GitHub user
	GitHubUserID string `json:"github_auth_token"`

	// GitHubAuthToken is the user's GitHub authentication token
	GitHubAuthToken string `json:"github_auth_token"`
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
