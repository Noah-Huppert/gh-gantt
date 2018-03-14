package main

import (
	"context"
	"log"
	"os"

	"github.com/Noah-Huppert/gh-gantt/cache"
	"github.com/Noah-Huppert/gh-gantt/config"
	"github.com/Noah-Huppert/gh-gantt/gh"
	"github.com/Noah-Huppert/gh-gantt/redis"
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

	// Redis clients
	redisClient := redis.NewClient(cfg)
	redisCache := cache.NewClient(cfg)

	// Server
	logger.Printf("starting server")

	srv := server.NewServer(ctx, cfg, ghClient, redisClient, redisCache)

	err = srv.Start()
	if err != nil {
		logger.Fatalf("error starting server: %s", err.Error())
	}
}
