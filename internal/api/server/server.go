package server

import (
	"bff/config"
	"bff/internal/api/router"
	"fmt"
	"net/http"
)

// Start initializes and starts the HTTP server.
func Start(cfg *config.Config) error {
	// Setup the main router
	mainRouter, err := router.Setup(cfg)
	if err != nil {
		return err
	}

	// Start the server
	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.AppPort), mainRouter)
}
