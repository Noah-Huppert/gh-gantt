package auth

import (
	"net/http"
	"net/url"

	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
)

// AuthLoginHandler implements resp.ResponderHandler by redirecting users to the GitHub OAuth login page
type AuthLoginHandler struct {
	// logger is used to display debug information
	logger golog.Logger

	// cfg is the application configuration
	cfg config.Config
}

// NewAuthLoginHandler creates a new AuthLoginHandler
func NewAuthLoginHandler(logger golog.Logger, cfg config.Config) AuthLoginHandler {
	return AuthLoginHandler{
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h AuthLoginHandler) Handle(r *http.Request) resp.Responder {
	// Construct redirect URL to GitHub OAuth page
	redirectURL, err := url.Parse("https://github.com/login/oauth/authorize")

	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error building GitHub OAuth redirect URL", err.Error())
	}

	queryParams := redirectURL.Query()
	queryParams.Set("client_id", h.cfg.GithubClientID)

	redirectURL.RawQuery = queryParams.Encode()

	return resp.NewRedirectResponder(redirectURL.String(), false)
}
