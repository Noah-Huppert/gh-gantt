package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/rpc"
	"github.com/Noah-Huppert/gh-gantt/server"
)

// logger used to output program information
var logger *log.Logger = log.New(os.Stdout, "main: ", 0)

func main() {
	ctx, ctxCancel := context.WithCancel(context.Background())

	// Setup sigterm handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		select {
		case <-sigChan:
			logger.Println("caught interrupt, attempting to " +
				"shutdown gracefully")
			ctxCancel()
		}
	}()

	// Configuration
	logger.Println("loading configuration")

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("error loading configuration: %s\n", err.Error())
	}

	// GitHub setup
	ghClient := gh.NewClient(ctx, cfg)

	// Status channel setup

	// statusOKChan receives the names of servers which have exited
	// successfully
	statusOKChan := make(chan string, 2)

	// statusErrChan receives errors from the rpc or http servers
	statusErrChan := make(chan error, 2)

	// GRPC Server
	logger.Printf("starting grpc server on :%d\n", cfg.RPC.Port)

	rpc.Start(ctx, statusOKChan, statusErrChan, cfg)

	// HTTP Server
	logger.Printf("starting http server on :%d\n", cfg.HTTP.Port)

	srv := server.NewServer(ctx, cfg, ghClient)

	srv.Start(statusOKChan, statusErrChan)

	// Wait for servers to finish
	shutdownOK := true
	for i := 0; i < 2; i++ {
		select {
		case system := <-statusOKChan:
			logger.Printf("%s successfully shutdown\n",
				system)

		case err = <-statusErrChan:
			logger.Println(err.Error())
			shutdownOK = false
		}
	}

	if shutdownOK {
		logger.Println("successfully shutdown")
	} else {
		logger.Fatalf("failed to shutdown")
	}
}
