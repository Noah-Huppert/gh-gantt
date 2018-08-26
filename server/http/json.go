package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponder implements the Responder interface by sending a JSON response to the client
type JSONResponder struct {
	// Data is the JSON data to respond with
	Data interface{}

	// Status is the HTTP status code to respond with, defaults to 200
	Status int
}

// NewJSONResponder creates a new JSONResponder
func NewJSONResponder(data interface{}, status int) JSONResponder {
	return JSONResponder{
		Data:   data,
		Status: status,
	}
}

// Respond implements Responder.Response
func (r JSONResponder) Response(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)

	err := encoder.Encode(r.Data)

	if err != nil {
		panic(fmt.Errorf("error encoding response into JSON: %s", err.Error()))
	}

	w.WriteHeader(r.Status)
}

// ServeHTTP implements http.Handler
func (h JSONResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Call ResultHandler
	result, err := h.resultHandler(r)

	// JSON encode result and send as response
	encoder := json.NewEncoder(w)

	if err != nil {
		// If error send JSON encoded error as response
		resp := map[string]string{
			"error": err.Error(),
		}

		err = encoder.Encode(resp)

		// If failed to JSON encode
		if err != nil {
			h.sendJSONEncodingErr(w, err)
		} else {
			// If success, write custom error status
			w.WriteHeader(err.Status())
		}
	} else {
		// If no error send result
		err = encoder.Encode(result)

		err = encoder.Encode(result)

		// If failed to JSON encode
		if err != nil {
			h.sendJSONEncodingErr(w, err)
		}
	}
}

// sendJSONEncodingErr sends an error related to JSON encoding to the user. If in the process of sending this error
// another error occurs a pre-canned string which is valid JSON is sent as the response
func (h JSONResultHandler) sendJSONEncodingErr(w http.ResponseWriter, err error) {
	w.WriterHeader(http.StatusInternalServerError)

	// Try to send actual error in new JSON encoded object
	encoder := json.Encoder(w)

	resp := map[string]string{
		"error": fmt.Sprintf("failed to encode response into JSON: %s", err.Error()),
	}

	err := encoder.Encode(resp)

	if err != nil {
		// If an error occurs encoding this object, send pre-canned string which is valid JSON
		fmt.Fprint(w, "{\"error\": \"an unknown error occurred while encoding the HTTP response into JSON\"}")
	}
}
