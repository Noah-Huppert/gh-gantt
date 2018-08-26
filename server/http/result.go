package http

import (
	"net/http"
)

// ResponderHandler handles HTTP requests using a handler method which returns a Responder.
// This Responder can then be invoked by a standard http.Handler to respond to an HTTP request.
type ResponderHandler interface {
	// Handle handles an HTTP request and returns a Responder which will be used to respond to the HTTP request
	Handle(r *http.Request) Responder
}

// ResponderHandlerWrapper calls a ResponderHandler in its http.Handler.ServeHTTP method. So a ResponderHandler can be
// used as a standard http.Handler
type ResponderHandlerWrapper struct {
	// responderHandler is the ResponderHandler to call in the standard http.Handler.ServeHTTP method
	responderHandler ResponderHandler
}

// ServeHTTP implements http.Handler.ServeHTTP
func (w ResponderHandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responder := w.responderHandler.Handle(r)

	responder.Respond(w, r)
}

// WrapResponderHandler wraps a ResponderHandler in a ResponderHandlerWrapper so it can be called as a
// standard http.Handler
func WrapResponderHandler(responderHandler ResponderHandler) ResponderHandlerWrapper {
	return ResponderHandlerWrapper{
		responderHandler: responderHandler,
	}
}
