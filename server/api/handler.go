package api

import (
	"github.com/gorilla/mux"
)

// SetupRouter registers API routes with the provided router
func SetupRouter(router *mux.Router) {
	router.Handle("/healthz", NewHealthCheckHandler()).Methods("GET")
}
