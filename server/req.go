package server

import (
	"fmt"
	"net/http"
)

// RequireBodyFields ensures the specified fields are present in the body. If
// they are not an error response is sent.
//
// A bool is returned to indicate if the required body fields are all present.
func RequireBodyFields(w http.ResponseWriter, r *http.Request, fields []string) bool {
	if err := r.ParseForm(); err != nil {
		WriteErr(w, 500, fmt.Errorf("error parsing request "+
			"body: %s", err.Error()))
		return false
	}

	// Check each field exists
	missing := []string{}

	for _, field := range fields {
		if _, ok := r.Form[field]; !ok {
			missing = append(missing, field)
		}
	}

	// If any fields missing
	if len(missing) > 0 {
		WriteErr(w, 400, fmt.Errorf("[%s] fields missing from body",
			missing))
		return false
	}

	// All present
	return true
}
