package resp

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

// Respond implements Responder.Respond
func (r JSONResponder) Respond(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)

	err := encoder.Encode(r.Data)

	if err != nil {
		panic(fmt.Errorf("error encoding response into JSON: %s", err.Error()))
	}
}
