// Package handler provides HTTP request handlers for the CoCDB REST API.
// It includes standardized response formatting for consistent API output.
package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// APIResponse is the standard envelope for all successful API responses.
type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Meta   *Meta       `json:"meta,omitempty"`
}

// Meta provides additional response metadata.
type Meta struct {
	Cached bool `json:"cached"`
}

// ErrorResponse is the standard envelope for error responses.
type ErrorResponse struct {
	Status string    `json:"status"`
	Error  ErrorInfo `json:"error"`
}

// ErrorInfo contains details about an API error.
type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// writeJSON marshals data to JSON and writes it with the given HTTP status code.
func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("failed to encode JSON response", "error", err)
	}
}

// Success sends a standardized successful response with optional metadata.
func Success(w http.ResponseWriter, data interface{}, meta *Meta) {
	writeJSON(w, http.StatusOK, APIResponse{
		Status: "success",
		Data:   data,
		Meta:   meta,
	})
}

// Error sends a standardized error response with the given HTTP status and message.
func Error(w http.ResponseWriter, code int, message string) {
	writeJSON(w, code, ErrorResponse{
		Status: "error",
		Error:  ErrorInfo{Code: code, Message: message},
	})
}

// NotFound sends a 404 error response.
func NotFound(w http.ResponseWriter, message string) {
	Error(w, http.StatusNotFound, message)
}

// InternalError logs the error and sends a 500 error response.
func InternalError(w http.ResponseWriter, message string) {
	slog.Error("internal server error", "message", message)
	Error(w, http.StatusInternalServerError, message)
}
