package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ParseBody parses a request's body from JSON form into a map.
// An error is returned if one occurs
func ParseBody(r *http.Request, body interface{}) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		return fmt.Errorf("error decoding JSON body: %s",
			err.Error())
	}

	return nil
}

// RequireBodyFields ensures the specified fields are present in the body. If
// they are not an error response is sent.
//
// Returns a bool indicating if required fields are present.
func RequireBodyFields(w http.ResponseWriter, r *http.Request,
	fields []string) bool {

	body := map[string]interface{}{}
	err := ParseBody(r, &body)
	if err != nil {
		WriteErr(w, 500, fmt.Errorf("error parsing body: %s",
			err.Error()))
		return false
	}

	// Check each field exists
	missing := []string{}

	for _, field := range fields {
		if _, ok := body[field]; !ok {
			missing = append(missing, field)
		}
	}

	// If any fields missing
	if len(missing) > 0 {
		plurality := ""
		if len(missing) > 1 {
			plurality = "s"
		}

		WriteErr(w, 400, fmt.Errorf("%s field%s missing from body",
			missing, plurality))
		return false
	}

	// All present
	return true
}
