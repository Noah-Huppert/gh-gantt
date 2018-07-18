package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/google/go-github/github"

	"github.com/gorilla/mux"
)

// Server wraps the HTTP server functionality
type Server struct {
	ctx      context.Context
	cfg      *config.Config
	ghClient *github.Client
}

// NewServer creates a new Server instance.
func NewServer(ctx context.Context, cfg *config.Config,
	ghClient *github.Client) Server {

	return Server{
		ctx:      ctx,
		cfg:      cfg,
		ghClient: ghClient,
	}
}

// Registerables returns a slice of Registerable interface to be registered
// with a http.ServeMux.
func (s Server) Registerables() []Registerable {

	return []Registerable{
		NewNotFoundHandler(),
		NewIssuesEndpoint(s.ctx, s.cfg, s.ghClient),
		NewStaticFiles(),
	}
}

// Routes returns a mux.Router with all the server route handlers.
func (s Server) Routes() *mux.Router {
	router := mux.NewRouter()

	regs := s.Registerables()
	for _, reg := range regs {
		reg.Register(router)
	}

	return router
}

func (s Server) Start(statusOKChan chan<- string,
	statusErrChan chan<- error) {

	router := s.Routes()

	// Setup graceful stop handler
	portStr := fmt.Sprintf(":%d", s.cfg.HTTP.Port)

	httpServer := &http.Server{
		Addr:    portStr,
		Handler: router,
	}

	// startFailedChan will receive a message if the http server failed to
	// start. Content of message does not matter, any message received
	// indicates a start failure.
	//
	// Used by the graceful exit go routine to exit when the server does
	// not start.
	startFailedChan := make(chan bool, 1)

	go func() {
		select {
		case <-s.ctx.Done():
			err := httpServer.Shutdown(nil)

			if err != nil {
				statusErrChan <- fmt.Errorf("error "+
					"shutting down http server: %s",
					err.Error())
			} else {
				statusOKChan <- "http server"
			}
		case <-startFailedChan:
			return
		}
	}()

	// Serve http content
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		statusErrChan <- fmt.Errorf("error starting http server: %s",
			err.Error())
		startFailedChan <- true
	}
}
