package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StaticDir is the directory to serve static files from
const StaticDir string = "./static"

// StaticURL is the URL to register the static file handler at
const StaticURL string = "/"

// StaticFiles registers a static file file handlers for the front end
type StaticFiles struct {
	// BaseURL is the URL to serve static files from
	BaseURL string

	// Dir is the directory to serve files out of
	Dir string
}

// NewStaticFiles creates a new StaticFiles instance with the default values
func NewStaticFiles() StaticFiles {
	return StaticFiles{
		BaseURL: StaticURL,
		Dir:     StaticDir,
	}
}

// Register implements Registerable.Register
func (s StaticFiles) Register(router *mux.Router) {
	router.PathPrefix(s.BaseURL).Handler(http.FileServer(http.Dir(s.Dir)))
}
