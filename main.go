package main

import (
	"context"
	"log"
	"os"

	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/server"
)

// logger used to output program information
var logger *log.Logger = log.New(os.Stdout, "main: ", 0)

func main() {
	// Configuration
	logger.Printf("loading configuration")

	cfg, err := config.Load()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// GitHub client
	ctx := context.Background()

	ghClient := gh.NewClient(ctx, cfg)

	// Server
	logger.Printf("starting server on :%d", cfg.HTTP.Port)

	srv := server.NewServer(ctx, cfg, ghClient)

	err = srv.Start()
	if err != nil {
		logger.Fatalf("error starting server: %s", err.Error())
	}
}
