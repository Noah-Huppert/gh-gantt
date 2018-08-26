package http

import (
	"net/http"
)

// NestedHandle registers an http.Handler under a sub path in an http.ServeMux.
// This allows a handler to be created without the knowledge of what sub path it will be registered under.
//
// Ex: Handler responds to /foo route. With this method it could be registered under /api/foo with no changes to the
// handlers code.
func NestedHandle(parent *http.ServeMux, path string, child http.Handler) {
	// Make versions of path with and without trailing slash
	lenPath := len(path)

	withTrailingSlash := path
	if withTrailingSlash[lenPath-1] != '/' {
		withTrailingSlash += "/"
	}

	withoutTrailingSlash := path
	if withoutTrailingSlash[lenPath-1] == '/' {
		withoutTrailingSlash = withoutTrailingSlash[:lenPath-1]
	}

	// Register nested mux
	parent.Handle(withTrailingSlash, http.StripPrefix(withoutTrailingSlash, child))
}
