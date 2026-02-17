package handler

import (
	"net/http"
	"os"
)

// FaviconHandler serves a favicon from a file path on disk.
// If no favicon file exists yet, it returns 204 No Content.
type FaviconHandler struct {
	filePath string
}

// NewFaviconHandler creates a handler that serves the favicon at the given path.
// Place your favicon.ico in the static/ directory to enable it.
func NewFaviconHandler(filePath string) *FaviconHandler {
	return &FaviconHandler{filePath: filePath}
}

// ServeHTTP serves the favicon file or returns 204 if none is configured.
// Route: GET /favicon.ico
func (h *FaviconHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat(h.filePath); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.ServeFile(w, r, h.filePath)
}
