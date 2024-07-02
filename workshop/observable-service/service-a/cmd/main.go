package main

import (
	"api"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-logr/stdr"
	"go.opentelemetry.io/otel"
)

func main() {
	// Default logging to stdout + OpenTelemetry
	logger := stdr.New(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile))
	otel.SetLogger(logger)

	// Initialize
	api.SetupOTelSDK(context.Background())

	handler := api.SetupHandler()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Info("error" + err.Error())
		}
	}()

	<-done
	logger.Info("Shutting down server...")

	// Close all database connection etc....
	/*
	  Insert Code
	*/
	logger.Info("Waiting for 5 seconds before initiating shutdown...")
	time.Sleep(5 * time.Second)

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Info("Error while shutting down Server. Initiating force shutdown...", "Error", err.Error())
	} else {
		logger.Info("Server shutdown successfully", "error", "")
	}
}
