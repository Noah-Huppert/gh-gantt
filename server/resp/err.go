package resp

import (
	"errors"
	"net/http"

	"github.com/Noah-Huppert/golog"
)

// ErrorResponder implements the Responder interface by sending an error to the client
type ErrorResponder struct {
	// logger is used to display the DetailError
	logger golog.Logger

	// Status is the HTTP status code to set when sending the error
	Status int

	// PublicError is the error to return to the user
	PublicError error

	// DetailError provides more details error information which will not be shown to the end user. Can be nil
	DetailError error
}

// NewErrorResponder creates an ErrorResponder
func NewErrorResponder(logger golog.Logger, status int, pubErr error, detErr error) ErrorResponder {
	return ErrorResponder{
		logger:      logger,
		Status:      status,
		PublicError: pubErr,
		DetailError: detErr,
	}
}

// NewStrErrorResponder creates an ErrorResponder. An error for PublicError and DetailError will be created from the
// error string arguments.
func NewStrErrorResponder(logger golog.Logger, status int, pubErrStr, detErrStr string) ErrorResponder {
	var detErr error = nil
	if len(detErrStr) > 0 {
		detErr = errors.New(detErrStr)
	}

	return ErrorResponder{
		logger:      logger,
		Status:      status,
		PublicError: errors.New(pubErrStr),
		DetailError: detErr,
	}
}

// Respond implements Responder.Respond
func (e ErrorResponder) Respond(w http.ResponseWriter, r *http.Request) {
	if e.PublicError == nil {
		panic("ErrorResponder.PublicError can not be nil")
	}

	if e.DetailError != nil {
		e.logger.Errorf("%s: %s", e.PublicError.Error(), e.DetailError.Error())
	} else {
		e.logger.Error(e.PublicError.Error())
	}

	resp := map[string]string{
		"error": e.PublicError.Error(),
	}

	jsonResponder := NewJSONResponder(resp, e.Status)
	jsonResponder.Respond(w, r)
}
