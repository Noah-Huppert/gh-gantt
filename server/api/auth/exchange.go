package auth

import (
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/resp"
)

// AuthExchangeHandler implements resp.ResponderHandler by exchanging a GitHub OAuth temporary code for a longer lived
// GitHub auth token
type AuthExchangeHandler struct{}

// NewAuthExchangeHandler creates a new AuthExchangeHandler
func NewAuthExchangeHandler() AuthExchangeHandler {
	return AuthExchangeHandler{}
}

// Handle implements resp.ResponderHandler.Handle
func (h AuthExchangeHandler) Handle(r *http.Request) resp.Responder {
	return resp.NewJSONResponder("ok", http.StatusOK)
}
