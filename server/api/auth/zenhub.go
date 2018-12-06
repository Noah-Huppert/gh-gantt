package auth

import (
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/auth"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/req"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
)

// ZenHubAppendHandler implements resp.ResponderHandler by appending a provided ZenHub authentication token to a
// provided API authentication token
type ZenHubAppendHandler struct {
	// logger prints debug information
	logger golog.Logger

	// cfg is application configuration
	cfg config.Config
}

// ZenHubAppendRequest is the format of a ZenHub append request
type ZenHubAppendRequest struct {
	// AuthToken is an existing API authentication token
	AuthToken string `json:"auth_token" validate:"nonzero"`

	// ZenHubAuthToken is the ZenHub authentication token to append to the AuthToken
	ZenHubAuthToken string `json:"zenhub_auth_token" validate:"nonzero"`
}

// NewZenHubAppendHandler creates a new ZenHubAppendHandler
func NewZenHubAppendHandler(logger golog.Logger, cfg config.Config) ZenHubAppendHandler {
	return ZenHubAppendHandler{
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h ZenHubAppendHandler) Handle(r *http.Request) resp.Responder {
	// Decode body
	var request ZenHubAppendRequest

	errResp := req.DecodeValidatedJSON(h.logger, r, &request)
	if errResp != nil {
		return errResp
	}

	// Parse provided API auth token
	authToken := &auth.AuthToken{}
	err := authToken.Decode(request.AuthToken, h.cfg.SigningSecret)

	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error parsing provided API authentication token", err.Error())
	}

	// Append ZenHub auth token
	authToken.ZenHubAuthToken = request.ZenHubAuthToken

	// Encode new auth token
	authTokenStr, err := authToken.Encode(h.cfg.SigningSecret)

	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error encoding new API authentication token", err.Error())
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"auth_token": authTokenStr,
	}, http.StatusOK)
}
