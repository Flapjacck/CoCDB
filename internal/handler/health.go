package handler

import (
	"net/http"
	"time"
)

// HealthHandler provides a health-check endpoint for monitoring and load balancers.
type HealthHandler struct {
	startTime time.Time
	version   string
}

// NewHealthHandler creates a HealthHandler that tracks server uptime.
func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		version:   version,
	}
}

// healthResponse is the structure returned by the health endpoint.
type healthResponse struct {
	Status    string `json:"status"`
	Version   string `json:"version"`
	Uptime    string `json:"uptime"`
	Timestamp string `json:"timestamp"`
}

// ServeHTTP responds with current server health information.
// Route: GET /health
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Status:    "healthy",
		Version:   h.version,
		Uptime:    time.Since(h.startTime).Round(time.Second).String(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
