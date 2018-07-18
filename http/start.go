package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Noah-Huppert/gh-gantt/config"
)

// StaticDir is the directory to serve static files from
const staticDir string = "./static"

// Start serving http content
func Start(ctx context.Context, statusOKChan chan<- string,
	statusErrChan chan<- error, cfg *config.Config) {

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
