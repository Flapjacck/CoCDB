package handler

import (
	"net/http"
	"time"
)

// HealthHandler provides a health-check endpoint for monitoring and load balancers.
type HealthHandler struct {
	startTime time.Time
}

// NewHealthHandler creates a HealthHandler that tracks server uptime.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
	}
}

// healthResponse is the structure returned by the health endpoint.
type healthResponse struct {
	Status    string `json:"status"`
	Uptime    string `json:"uptime"`
	Timestamp string `json:"timestamp"`
}

// ServeHTTP responds with current server health information.
// Route: GET /health
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{
		Status:    "healthy",
		Uptime:    time.Since(h.startTime).Round(time.Second).String(),
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}
