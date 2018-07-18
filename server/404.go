package server

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// NotFoundHandler returns a 404 not found API response
type NotFoundHandler struct{}

// NewNotFoundHandler creates a new NotFoundHandler instance
func NewNotFoundHandler() NotFoundHandler {
	return NotFoundHandler{}
}

// Register implements Registerable.Register
func (h NotFoundHandler) Register(router *mux.Router) {
	router.NotFoundHandler = h
}

func (h NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	WriteErr(w, 404, errors.New("not found"))
}
