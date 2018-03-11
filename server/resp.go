package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSON sends a JSON HTTP response.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	// Set response data type headers
	w.Header().Set("Content-Type", "application/json")

	// Set response HTTP status code
	w.WriteHeader(status)

	// Write data
	err := json.NewEncoder(w).Encode(data)

	// If error return manually encoded JSON string
	if err != nil {
		fmt.Fprintf(w, "{\"error\": \"error marshalling data into "+
			"JSON: %s\"}", err.Error())
	}
}

// WriteErr sends an error HTTP response.
func WriteErr(w http.ResponseWriter, status int, err error) {
	resp := map[string]string{
		"error": err.Error(),
	}

	WriteJSON(w, status, resp)
}
