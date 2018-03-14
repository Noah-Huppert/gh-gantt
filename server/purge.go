package server

import (
	"fmt"
	"net/http"
)

// PurgePath is the path to register Purge handlers at
const PurgePath string = "/api/cache/purge"

// PurgeEndpoint implements HTTP handlers for the purge endpoint
type PurgeEndpoint struct {
	// BasePath is the URL HTTP Purge handlers will be registered at
	BasePath string
}

// NewPurgeEndpoint creates a new PurgeEndpoint instance
func NewPurgeEndpoint() *PurgeEndpoint {
	return &PurgeEndpoint{
		BasePath: PurgePath,
	}
}

// Register implements Registerable.Register
func (p PurgeEndpoint) Register(mux *http.ServeMux) {
	mux.HandleFunc(p.BasePath, p.Post)
}

// Post handles purge endpoint post requests
func (p PurgeEndpoint) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "purge")
}
