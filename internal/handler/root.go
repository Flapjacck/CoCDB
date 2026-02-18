package handler

import "net/http"

// rootResponse describes the API at the root endpoint.
type rootResponse struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Endpoints   map[string]string `json:"endpoints"`
}

// RootHandler returns a handler that displays API information and available endpoints.
// Route: GET /
func RootHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, rootResponse{
			Name:        "CoCDB API",
			Description: "Clash of Clans Database REST API",
			Endpoints: map[string]string{
				"health":    "/health",
				"buildings": "/api/buildings",
				"troops":    "/api/troops",
			},
		})
	}
}
