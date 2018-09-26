package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/Noah-Huppert/gh-gantt/server/config"
	"github.com/Noah-Huppert/gh-gantt/server/encryption"
	"github.com/Noah-Huppert/gh-gantt/server/libdb"
	"github.com/Noah-Huppert/gh-gantt/server/serve"

	"github.com/Noah-Huppert/golog"
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

	cipherText, err := encryption.AESEncrypt([]byte(cfg.ThirdPartyAuthTokenEncryptionSecret), []byte("to encrypt"))
	if err != nil {
		logger.Fatalf("error encrypting: %s", err.Error())
	}
	logger.Debugf("encrypted: %s", string(cipherText))

	plainText, err := encryption.AESDecrypt([]byte(cfg.ThirdPartyAuthTokenEncryptionSecret), cipherText)
	if err != nil {
		logger.Fatalf("error decrypting: %s", err.Error())
	}
	logger.Debugf("decrypted: %s", string(plainText))
	return

	// Cancel context on interrupt signal
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	go func() {
		<-interruptChan
		ctxCancel()
	}()

	// Connect to database
	db, err := libdb.ConnectX(cfg.DBConfig)
	if err != nil {
		logger.Fatalf("failed to connect to database: %s", err.Error())
	}

	// Start HTTP server
	server := serve.NewServer(ctx, *cfg, db, logger)

	err = server.Serve()
	if err != nil {
		logger.Fatalf("error running HTTP server: %s", err.Error())
	}

	logger.Info("done")
}
