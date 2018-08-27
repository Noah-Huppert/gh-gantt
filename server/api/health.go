package api

import (
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/resp"
)

// HealthCheckHandler implements http.ResponderHandler for the health check endpoint
type HealthCheckHandler struct{}

// NewHealthCheckHandler creates a new HealthCheckHandler
func NewHealthCheckHandler() HealthCheckHandler {
	return HealthCheckHandler{}
}

// Handle implements resp.ResponderHandler.Handle
func (HealthCheckHandler) Handle(r *http.Request) resp.Responder {
	respBody := map[string]bool{
		"ok": true,
	}

	return resp.NewJSONResponder(respBody, http.StatusOK)
}
