package api

import (
	"fmt"
	"net/http"
)

// NewAPIHandler creates a new http.Handler to serve API requests
func NewAPIHandler() http.Handler {
	// Create mux
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	return mux
}
