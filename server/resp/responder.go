package resp

import (
	"net/http"
)

// Responder provides an extendible way to respond to HTTP requests
type Responder interface {
	// Respond sends a response to an HTTP request
	Respond(w http.ResponseWriter, r *http.Request)
}
