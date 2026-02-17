// CoCDB

package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/flapjacck/CoCDB/internal/config"
	"github.com/flapjacck/CoCDB/internal/router"
)

func main() {
	// Load configuration from environment variables.
	cfg := config.Load()

	// Set up structured logging based on environment.
	initLogger(cfg)

	slog.Info("starting CoCDB API server",
		"port", cfg.Port,
		"environment", cfg.Environment,
		"version", cfg.Version,
	)

	// Build the HTTP router with all routes and middleware.
	r := router.New(cfg)

	// Configure the HTTP server with timeouts for production resilience.
	srv := &http.Server{
		Addr:         cfg.Addr(),
		Handler:      r,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in a goroutine so we can listen for shutdown signals.
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	slog.Info("server is ready and accepting connections", "addr", cfg.Addr())

	// Block until we receive a termination signal (Ctrl+C, SIGTERM, etc.).
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	slog.Info("received shutdown signal", "signal", sig.String())

	// Give active connections up to 30 seconds to finish.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped gracefully")
}

// initLogger configures the global slog logger.
// Production uses JSON output for structured log aggregation.
// Development uses human-readable text format.
func initLogger(cfg *config.Config) {
	var level slog.Level
	switch cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{Level: level}

	var h slog.Handler
	if cfg.IsProd() {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(h))
}
