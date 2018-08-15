package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/status"
)

// staticDir is the directory to serve static files from
const staticDir string = "./frontend/dist"

// systemName is the name used to identify the http server to other pieces of
// the application
const systemName string = "http server"

// Start serving http content
func Start(ctx context.Context, statusChan chan<- status.StatusMsg,
	cfg *config.Config) {

	// Setup graceful stop handler
	portStr := fmt.Sprintf(":%d", cfg.HTTP.Port)

	httpServer := &http.Server{
		Addr:    portStr,
		Handler: http.FileServer(http.Dir(staticDir)),
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
		case <-ctx.Done():
			err := httpServer.Shutdown(ctx)

			if err != nil {
				statusChan <- status.StatusMsg{
					System: systemName,
					Err: fmt.Errorf("error shutting "+
						"down http server: %s",
						err.Error()),
				}
			} else {
				statusChan <- status.StatusMsg{
					System: systemName,
				}
			}
		case <-startFailedChan:
			return
		}
	}()

	// Serve http content
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		statusChan <- status.StatusMsg{
			System: systemName,
			Err: fmt.Errorf("error starting http server: %s",
				err.Error()),
		}
		startFailedChan <- true
	}
}
