package server

import (
	"net/http"
)

// Registerable is an interface used to register HTTP server handlers with a
// http.ServerMux.
type Registerable interface {
	// Register adds a HTTP handler to a http.ServerMux
	Register(mux *http.ServeMux)
}
