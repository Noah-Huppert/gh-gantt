package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/http"
	"github.com/Noah-Huppert/gh-gantt/rpc"
	"github.com/Noah-Huppert/gh-gantt/status"
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
	//ghClient := gh.NewClient(ctx, cfg)

	// Status channel setup

	// statusChan receives system component status messages
	statusChan := make(chan status.StatusMsg, 2)

	// GRPC Server
	logger.Printf("starting grpc server on :%d\n", cfg.RPC.Port)

	rpc.Start(ctx, statusChan, cfg)

	// HTTP Server
	logger.Printf("starting http server on :%d\n", cfg.HTTP.Port)

	http.Start(ctx, statusChan, cfg)

	// Wait for servers to finish
	shutdownOK := true
	for i := 0; i < 2; i++ {
		msg := <-statusChan

		if msg.Err != nil {
			logger.Printf("error in %s: %s\n", msg.System,
				err.Error())
			shutdownOK = false
		} else {
			logger.Printf("%s successfully shutdown\n", msg.System)
		}
	}

	if shutdownOK {
		logger.Println("successfully shutdown")
	} else {
		logger.Fatalf("failed to shutdown")
	}
}
