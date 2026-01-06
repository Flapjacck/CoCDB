package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// GetDataHandler retrieves JSON data from the data directory based on the provided path
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	// Get the path parameter
	path := chi.URLParam(r, "*")

	// Validate the path to prevent directory traversal
	if strings.Contains(path, "..") {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	// Construct the file path
	filePath := fmt.Sprintf("data/%s.json", strings.TrimSuffix(path, "/"))

	// Read the JSON file
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Building/Troop not found", http.StatusNotFound)
		return
	}

	// Parse and validate JSON
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
