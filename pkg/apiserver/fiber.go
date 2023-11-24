package apiserver

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"webhook/pkg/configs"

	"github.com/gofiber/fiber/v2"
)

// StartFiberWithGracefulShutdown function for starting server with a graceful shutdown.
func StartFiberWithGracefulShutdown(a *fiber.App, config *configs.ServerConfig) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnURL, _ := connectionURLBuilder(config)

	// Run server.
	if err := a.Listen(fiberConnURL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// ConnectionURLBuilder func for building URL connection.
func connectionURLBuilder(config *configs.ServerConfig) (string, error) {

	// URL for Fiber connection.
	url := fmt.Sprintf(
		"%s:%v",
		config.Host,
		config.Port,
	)

	// Return connection URL.
	return url, nil
}
