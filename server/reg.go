package server

import (
	"github.com/gorilla/mux"
)

// Registerable is an interface used to register HTTP server handlers with a
// http.ServerMux.
type Registerable interface {
	// Register adds a HTTP handler to a mux.Router
	Register(router *mux.Router)
}
