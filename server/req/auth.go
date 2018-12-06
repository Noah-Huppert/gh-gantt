package req

import (
	"net/http"
	"strings"

	"github.com/Noah-Huppert/gh-gantt/server/auth"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
)

// CheckAuthToken ensures an API authentication token is provided via the "Authorization" header. If present the auth
// token will be returned. Otherwise a response which sends an error will be returned.
func CheckAuthToken(logger golog.Logger, cfg config.Config, r *http.Request) (*auth.AuthToken, resp.Responder) {
	// Get from header
	authorization := r.Header.Get("Authorization")

	if len(authorization) == 0 {
		return nil, resp.NewStrErrorResponder(logger, http.StatusUnauthorized,
			"no authentication token provided", "authorization header empty")
	}

	// Parse
	authorizationParts := strings.Split(authorization, " ")

	if len(authorizationParts) != 2 {
		return nil, resp.NewStrErrorResponder(logger, http.StatusUnauthorized,
			"authorization header not in correct format, must be: token <TOKEN>", "")
	}

	authTokenStr := authorizationParts[1]

	authToken := &auth.AuthToken{}
	err := authToken.Decode(authTokenStr, cfg.SigningSecret)

	if err != nil {
		return nil, resp.NewStrErrorResponder(logger, http.StatusInternalServerError,
			"error parsing provided API authentication token", err.Error())
	}

	return authToken, nil
}
