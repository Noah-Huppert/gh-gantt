package auth

import (
	"encoder/json"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/config"
)

// ExchangeGitHubCodeURL is the GitHub API URL used to exchange a GitHub code for a token
const exchangeGitHubCodeURL string = "https://github.com/login/oauth/access_token"

// ExchangeGitHubCodeReq holds the parameters required to make a GitHub API request to exchange a temporary GitHub
// code for a longer lasting GitHub token
type ExchangeGitHubCodeReq struct {
	// ClientID is the GitHub app's Client ID
	ClientID string `json:"client_id"`

	// ClientSecret is the GitHub app's client secret
	ClientSecret string `json:"client_secret"`

	// Code is the temporary GitHub authentication code
	Code string `json:"code"`

	// State is a value used to prevent cross site request forgery
	State string `json:"state"`
}

// exchangeGitHubCodeResp is the format of the exchange GitHub API response
type exchangeGitHubCodeResp struct {
	// AccessToken is the long lasting GitHub authentication token
	AccessToken string `json:"access_token"`
}

// NewExchangeGitHubCodeReq creates a new ExchangeGitHubCodeReq. Most of the fields can be filled by passing an
// config.Config object
func NewExchangeGitHubCodeReq(cfg config.Config, code, state string) ExchangeGitHubCodeReq {
	return ExchangeGitHubCodeReq{
		ClientID:     cfg.GitHubClientID,
		ClientSecret: cfg.GitHubClientSecret,
		Code:         code,
		State:        state,
	}
}

// Exchange exchanges a temporary GitHub code for a longer lasting GitHub token
func (r ExchangeGitHubCodeReq) Exchange() (string, error) {
	// Encode request body
	var body []byte
	encoder := json.NewEncoder(body)

	err := encoder.Encode(r)
	if err != nil {
		return "", fmt.Errorf("error JSON encoding request body: %s", err.Error())
	}

	// Make request
	resp, err := http.Post(exchangeGitHubCodeURL, "application/json", body)
	if err != nil {
		return "", fmt.Errorf("error making exchange GitHub API token request: %s", err.Error())
	}

	// Decode request
	var resp exchangeGitHubCodeResp
	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(resp)
	if err != nil {
		return "", fmt.Errorf("error decoding exchange GitHub API response body: %s", err.Error())
	}

	// Success
	return resp.AccessToken, nil
}