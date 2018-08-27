package resp

import (
	"net/http"
)

// RedirectResponder issues a redirect response
type RedirectResponder struct {
	// Location is the URL the client will be redirected to
	Location string

	// Permanent indicates if a redirect is permanent. If a redirect is permanent the 308 (permanent redirect) HTTP
	// status code is used. Otherwise the 307 (temporary redirect) status code is used.
	Permanent bool
}

// NewRedirectResponder creates a RedirectResponder
func NewRedirectResponder(location string, permanent bool) RedirectResponder {
	return RedirectResponder{
		Location:  location,
		Permanent: permanent,
	}
}

// Respond implements Responder.Respond
func (r RedirectResponder) Respond(w http.ResponseWriter, req *http.Request) {
	// Write status
	status := http.StatusTemporaryRedirect

	if r.Permanent {
		status = http.StatusPermanentRedirect
	}

	http.Redirect(w, req, r.Location, status)
}
