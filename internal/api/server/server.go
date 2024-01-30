package server

import (
	"bff/config"
	"bff/internal/api/router"
	"bff/log"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Start initializes and starts the HTTP server.
func Start(cfg *config.Config) error {
	// Setup the main router
	mainRouter, err := router.Setup(cfg)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	// Start the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.AppPort),
		Handler:      mainRouter,
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 80 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		log.Logger.Infof("caught signal: %s", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		log.Logger.Infof("completing background tasksat addr: %s", srv.Addr)

		wg.Wait()

		shutdownError <- nil
	}()

	log.Logger.Infof("starting background tasks at addr: %s", srv.Addr)

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	log.Logger.Infof("stopped server at address: %s", srv.Addr)

	return nil
}
