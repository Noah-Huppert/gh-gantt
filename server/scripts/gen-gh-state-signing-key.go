package main

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/Noah-Huppert/golog"
	"golang.org/x/crypto/ed25519"
)

const StateSigningKeyEnvKey = "APP_GH_STATE_SIGNING_KEY"

func main() {
	// Setup logger
	logger := golog.NewStdLogger("gen-gh-state-signing-key")

	// Generate key
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		logger.Fatalf("failed to generate key: %s", err.Error())
	}

	logger.Info("Generated key")

	// Write key to .env file
	f, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatalf("failed to open .env file to save key: %s", err.Error())
	}

	writeStr := fmt.Sprintf("\nexport %s=\"%s\"", StateSigningKeyEnvKey, privateKey)
	_, err = f.WriteString(writeStr)
	if err != nil {
		logger.Fatalf("failed to write .env file: %s", err.Error())
	}

	err = f.Close()
	if err != nil {
		logger.Fatalf("failed to close .env file: %s", err.Error())
	}

	logger.Info("Saved to file")
}
