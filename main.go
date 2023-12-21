package main

import (
	"bff/config"
	"bff/internal/api/server"
	"bff/log"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Initialize logging
	if err := log.Initialize(); err != nil {
		return err
	}

	// Load the configuration
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		return err
	}

	// Start the server
	return server.Start(cfg)
}
