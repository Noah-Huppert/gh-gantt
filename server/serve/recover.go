package serve

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/Noah-Huppert/golog"
)

// RecoveryHandler wraps a http.Handler and recovers from panics
type RecoveryHandler struct {
	// rootHandler is the HTTP handler which should be used to handle requests. If this handler panics while handling
	// requests the RecoveryHandler will take over
	rootHandler http.Handler

	// logger is used to print information about a panic
	logger golog.Logger
}

// NewRecoveryHandler creates a new RecoveryHandler
func NewRecoveryHandler(rootHandler http.Handler, logger golog.Logger) RecoveryHandler {
	return RecoveryHandler{
		rootHandler: rootHandler,
		logger:      logger,
	}
}

// ServeHTTP implements http.Handler.ServeHTTP
func (r RecoveryHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Setup recovery
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		r.logger.Errorf("%s %s panicked: %s", req.Method, req.URL.String(), err)
		debug.PrintStack()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprint(w, "{\"error\": \"an internal server error occurred\"}")
	}()

	r.rootHandler.ServeHTTP(w, req)
}
