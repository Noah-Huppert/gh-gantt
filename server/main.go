package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/serve"

	"github.com/Noah-Huppert/golog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Setup context
	ctx, ctxCancel := context.WithCancel(context.Background())

	// Setup logger
	logger := golog.NewStdLogger("gh-gantt")

	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("error loading configuration: %s", err.Error())
	}

	// Cancel context on interrupt signal
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		<-interruptChan
		ctxCancel()
	}()

	// Connect to database
	sqlConnStr := fmt.Sprintf("host=%s port=%d dbname=%s user=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUsername)

	if len(cfg.DBPassword) > 0 {
		sqlConnStr += fmt.Sprintf("password=%s", cfg.DBPassword)
	}

	db, err := sqlx.Connect("postgres", sqlConnStr)
	if err != nil {
		logger.Fatalf("error connecting to the database: %s", err.Error())
	}

	// Start HTTP server
	server := serve.NewServer(ctx, *cfg, db, logger)

	err = server.Serve()
	if err != nil {
		logger.Fatalf("error running HTTP server: %s", err.Error())
	}

	logger.Info("done")
}
