package http

import (
	"fmt"
	"net/http"
)

// ErrorResponder implements the Responder interface by sending an error to the client
type ErrorResponder struct {
	// Error is the error to return to the user
	Error error

	// Status is the HTTP status code to set when sending the error
	Status int
}

// NewErrorResponder creates an ErrorResponder
func NewErrorResponder(err error, status int) ErrorResponder {
	return ErrorResponder{
		Error:  err,
		Status: status,
	}
}

// Respond implements Responder.Respond
func (e ErrorResponder) Respond(w http.ResponseWriter, r http.Request) {
	resp := map[string]string{
		"error": e.Error.Error(),
	}

	jsonResponder := NewJSONResponder(resp, e.Status)
	jsonResponder.Respond(w, r)
}
