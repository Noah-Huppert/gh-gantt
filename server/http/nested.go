package http

import (
	"net/http"
)

// NestedMux registers a child http.ServeMux under a sub path in a parent http.ServeMux
func NestedMux(parent *http.ServeMux, path string, child *http.ServeMux) {
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
