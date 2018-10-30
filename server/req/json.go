package req

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/server/resp"

	"github.com/Noah-Huppert/golog"
	"gopkg.in/validator.v2"
)

// DecodeJSON decodes a JSON formatted request body. Returns an ErrorResponder if an error occurs while decoding
// the body.
func DecodeJSON(logger golog.Logger, r *http.Request, dest interface{}) resp.Responder {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(dest)
	if err != nil {
		return resp.NewStrErrorResponder(logger, http.StatusInternalServerError, "error decoding body", err.Error())
	}

	return nil
}

// DecodeValidatedJSON decodes a JSON formatted request body and validates it. Returns an ErrorResponder if an error
// occurs while decoding or validating the body.
func DecodeValidatedJSON(logger golog.Logger, r *http.Request, dest interface{}) resp.Responder {
	// Decode JSON
	errResp := DecodeJSON(logger, r, dest)
	if errResp != nil {
		return errResp
	}

	// Validate
	err := validator.Validate(dest)
	if err != nil {
		return resp.NewStrErrorResponder(logger, http.StatusBadRequest,
			fmt.Sprintf("error validating body: %s", err.Error()), err.Error())
	}

	return nil
}
