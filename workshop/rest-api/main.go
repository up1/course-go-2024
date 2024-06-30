package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Working with structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	handler := setupHandler()

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
			log.Printf("error: %s\n", err)
		}
	}()

	<-done
	log.Println("Shutting down server...")

	// Close all database connection etc....
	/*
	  Insert Code
	*/
	log.Println("Waiting for 5 seconds before initiating shutdown...")
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

func setupHandler() *gin.Engine {
	router := gin.New()

	// This handler will match /health
	router.GET("/health", func(c *gin.Context) {
		log.Print("Calling /health endpoint")
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})
	return router
}
