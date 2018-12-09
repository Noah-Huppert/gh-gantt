package auth

import (
	"context"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/auth"
	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/libgh"
	"github.com/Noah-Huppert/gh-gantt/server/req"
	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
)

// AuthExchangeHandler implements resp.ResponderHandler by exchanging a GitHub OAuth temporary code for a longer lived
// GitHub auth token
type AuthExchangeHandler struct {
	// ctx is the context
	ctx context.Context

	// logger is used to output debug information
	logger golog.Logger

	// cfg is the application configuration
	cfg config.Config
}

// AuthExchangeRequest is the format for an auth exchange request
type AuthExchangeRequest struct {
	// State is the state parameter returned in the GitHub redirect, used to prevent cross site forgery
	State string `json:"state" validate:"nonzero"`

	// Code is the temporary authentication code GitHub provides which can be swapped out for a longer living auth token
	Code string `json:"code" validate:"nonzero"`
}

// NewAuthExchangeHandler creates a new AuthExchangeHandler
func NewAuthExchangeHandler(ctx context.Context, logger golog.Logger, cfg config.Config) AuthExchangeHandler {
	return AuthExchangeHandler{
		ctx:    ctx,
		logger: logger,
		cfg:    cfg,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h AuthExchangeHandler) Handle(r *http.Request) resp.Responder {
	// Decode body
	var request AuthExchangeRequest

	errResp := req.DecodeValidatedJSON(h.logger, r, &request)
	if errResp != nil {
		return errResp
	}

	// Check state
	stateValid, err := auth.VerifyState(h.cfg.GetGHStateSigningPubKey(), request.State)
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError, "error validating state parameter",
			err.Error())
	}

	if !stateValid {
		return resp.NewStrErrorResponder(h.logger, http.StatusBadRequest, "invalid state", "stateValid=false")
	}

	// Exchange code
	exchangeReq := libgh.NewExchangeGitHubCodeRequest(h.cfg, request.Code, request.State)

	ghAuthToken, err := exchangeReq.Do()
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error exchanging code for GitHub access token", err.Error())
	}

	// Identify GitHub user
	identifyReq := libgh.NewIdentifyAuthTokenRequest(h.ctx, ghAuthToken)

	ghUserID, err := identifyReq.Do()
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error identifying user", err.Error())
	}

	// Create auth token
	authToken := auth.NewAuthToken(h.cfg.ServiceName, ghUserID, ghAuthToken, "")

	authTokenStr, err := authToken.Encode(h.cfg.SigningSecret)
	if err != nil {
		return resp.NewStrErrorResponder(h.logger, http.StatusInternalServerError,
			"error encoding authentication token", err.Error())
	}

	return resp.NewJSONResponder(map[string]interface{}{
		"auth_token": authTokenStr,
	}, http.StatusOK)
}
