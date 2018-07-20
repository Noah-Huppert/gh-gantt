package rpc

import (
	"context"
	"fmt"
	"net"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/repos"
	"github.com/Noah-Huppert/gh-gantt/status"

	"google.golang.org/grpc"
)

// systemName is the name used to identify the rpc server to other pieces of
// the application
const systemName string = "rpc server"

// Start the GRPC server.
func Start(ctx context.Context, statusChan chan<- status.StatusMsg,
	cfg *config.Config) {

	// Listen on rpc port
	grpcHostStr := fmt.Sprintf(":%d", cfg.RPC.Port)

	grpcListener, err := net.Listen("tcp", grpcHostStr)
	if err != nil {
		statusChan <- status.StatusMsg{
			System: systemName,
			Err: fmt.Errorf("error listening on grpc port: %s\n",
				err.Error()),
		}
		return
	}

	// Register services
	grpcServer := grpc.NewServer()

	repos.RegisterRepositoriesServer(grpcServer,
		repos.NewDefaultRepositoriesServer())

	// Setup graceful stop handler

	// startFailedChan will receive a message if the grpc server failed to
	// start. Content of message does not matter, any message received
	// indicates a start failure.
	//
	// Used by the graceful exit go routine to exit when the server does
	// not start.
	startFailedChan := make(chan bool, 1)

	go func() {
		select {
		case <-ctx.Done():
			grpcServer.GracefulStop()
			statusChan <- status.StatusMsg{
				System: systemName,
			}
		case <-startFailedChan:
			return
		}
	}()

	// Start serving requests
	go func() {
		err = grpcServer.Serve(grpcListener)
		if err != nil {
			statusChan <- status.StatusMsg{
				System: systemName,
				Err: fmt.Errorf("error serving rpc "+
					"requests: %s", err.Error()),
			}
			startFailedChan <- true
		}
	}()
}
