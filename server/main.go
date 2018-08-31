package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/libdb"
	"github.com/Noah-Huppert/gh-gantt/server/serve"

	"github.com/Noah-Huppert/golog"
	"github.com/jmoiron/sqlx"
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
	db, err := libdb.Connect(cfg.DBConfig)
	if err != nil {
		logger.Fatalf("error connecting to database: %s", err.Error())
	}

	dbx, err := sqlx.NewDb(db, "postgres")
	if err != nil {
		logger.Fatalf("error creating sqlx database instance from regular database instance: %s", err.Error())
	}

	// Start HTTP server
	server := serve.NewServer(ctx, *cfg, dbx, logger)

	err = server.Serve()
	if err != nil {
		logger.Fatalf("error running HTTP server: %s", err.Error())
	}

	logger.Info("done")
}
