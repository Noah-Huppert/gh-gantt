package auth

import (
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/jmoiron/sqlx"
)

// AuthExchangeHandler implements resp.ResponderHandler by exchanging a GitHub OAuth temporary code for a longer lived
// GitHub auth token
type AuthExchangeHandler struct {
	// db is a database connection
	db *sqlx.DB
}

// NewAuthExchangeHandler creates a new AuthExchangeHandler
func NewAuthExchangeHandler(db *sqlx.DB) AuthExchangeHandler {
	return AuthExchangeHandler{
		db: db,
	}
}

// Handle implements resp.ResponderHandler.Handle
func (h AuthExchangeHandler) Handle(r *http.Request) resp.Responder {
	return resp.NewJSONResponder("ok", http.StatusOK)
}
