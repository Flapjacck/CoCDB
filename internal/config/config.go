// Package config provides application configuration management.
// All settings are loaded from environment variables with sensible defaults,
// making it easy to configure for different environments (dev, staging, prod).
package config

import (
	"os"
	"strings"
	"time"
)

// Config holds all application configuration values.
type Config struct {
	// Server settings
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration

	// Application settings
	Environment string
	LogLevel    string
	DataDir     string
	Version     string

	// Cache settings
	CacheTTL time.Duration

	// CORS settings
	CORSOrigins []string
}

// Load reads configuration from environment variables, falling back to defaults.
//
// Environment variables:
//   - PORT: Server port (default: "3000")
//   - READ_TIMEOUT: HTTP read timeout (default: "10s")
//   - WRITE_TIMEOUT: HTTP write timeout (default: "10s")
//   - IDLE_TIMEOUT: HTTP idle timeout (default: "120s")
//   - ENVIRONMENT: Running environment (default: "development")
//   - LOG_LEVEL: Logging level — debug, info, warn, error (default: "info")
//   - DATA_DIR: Path to data directory (default: "data")
//   - APP_VERSION: Application version string (default: "1.0.0")
//   - CACHE_TTL: Cache time-to-live duration (default: "5m")
//   - CORS_ORIGINS: Comma-separated allowed origins (default: "*")
func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "3000"),
		ReadTimeout:  getDuration("READ_TIMEOUT", 10*time.Second),
		WriteTimeout: getDuration("WRITE_TIMEOUT", 10*time.Second),
		IdleTimeout:  getDuration("IDLE_TIMEOUT", 120*time.Second),
		Environment:  getEnv("ENVIRONMENT", "development"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		DataDir:      getEnv("DATA_DIR", "data"),
		Version:      getEnv("APP_VERSION", "1.0.0"),
		CacheTTL:     getDuration("CACHE_TTL", 5*time.Minute),
		CORSOrigins:  strings.Split(getEnv("CORS_ORIGINS", "*"), ","),
	}
}

// IsProd returns true when running in a production environment.
func (c *Config) IsProd() bool {
	return c.Environment == "production"
}

// Addr returns the formatted listen address (e.g., ":3000").
func (c *Config) Addr() string {
	return ":" + c.Port
}

// getEnv retrieves an environment variable or returns the fallback value.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// getDuration retrieves a duration from an environment variable.
// The value must be a valid Go duration string (e.g., "10s", "5m").
func getDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
