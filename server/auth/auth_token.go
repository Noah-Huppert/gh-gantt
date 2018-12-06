package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthToken is issued to a user to prove they are authenticated
type AuthToken struct {
	// Issuer is the name of the service which issued the token
	Issuer string `json:"iss"`

	// Audience is the name of the service who is meant to received the token
	Audience string `json:"aud"`

	// GitHubUserID is the ID of the authenticated GitHub user, used as the subject field in the JWT
	GitHubUserID string `json:"sub"`

	// GitHubAuthToken is the user's GitHub authentication token
	GitHubAuthToken string `json:"github_auth_token"`

	// ZenHubAuthToken is the user's ZenHub authentication token
	ZenHubAuthToken string `json:"zenhub_auth_token"`
}

// NewAuthToken creates a new AuthToken.
// The serviceName is used for the Issuer and Audience field
func NewAuthToken(serviceName, ghUserID, ghAuthToken, zhAuthToken string) AuthToken {
	return AuthToken{
		Issuer:          serviceName,
		Audience:        serviceName,
		GitHubUserID:    ghUserID,
		GitHubAuthToken: ghAuthToken,
		ZenHubAuthToken: zhAuthToken,
	}
}

// claims returns a map of JWT claims to encode
func (t AuthToken) claims() map[string]interface{} {
	return map[string]interface{}{
		"iss":               t.Issuer,
		"aud":               t.Audience,
		"sub":               t.GitHubUserID,
		"github_auth_token": t.GitHubAuthToken,
		"zenhub_auth_token": t.ZenHubAuthToken,
	}
}

// Encode an authentication token into a string
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

// Decode an authentication from a string
func (t *AuthToken) Decode(tokenStr, signingSecret string) error {
	// Parse
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method must be HS256, was: %s", token.Header["alg"])
		}

		return signingSecret, nil
	})

	if err != nil {
		return fmt.Errorf("error parsing JWT: %s", err.Error())
	}

	if !token.Valid {
		return errors.New("failed to verify token")
	}

	// Check claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("failed to convert claims to map")
	}

	requiredClaims := []string{"iss", "aud", "sub", "github_auth_token", "zenhub_auth_token"}
	missingClaims := []string{}

	for _, claimKey := range requiredClaims {
		if _, ok := claims[claimKey]; !ok {
			missingClaims = append(missingClaims, claimKey)
		}
	}

	if len(missingClaims) > 0 {
		return fmt.Errorf("missing claims: %s", strings.Join(missingClaims, ", "))
	}

	return nil
}
