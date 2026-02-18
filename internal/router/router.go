// Package router configures the HTTP router with all middleware and route handlers.
package router

import (
	"net/http"

	"github.com/flapjacck/CoCDB/internal/cache"
	"github.com/flapjacck/CoCDB/internal/config"
	"github.com/flapjacck/CoCDB/internal/data"
	"github.com/flapjacck/CoCDB/internal/handler"
	mw "github.com/flapjacck/CoCDB/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// New creates and configures a chi router with all API routes and middleware.
func New(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()

	// --- Global Middleware Stack ---
	r.Use(middleware.RequestID)          // Inject request ID into context
	r.Use(middleware.RealIP)             // Set RemoteAddr to X-Forwarded-For / X-Real-IP
	r.Use(mw.RequestLogger)              // Structured request logging
	r.Use(mw.SecurityHeaders)            // Security headers on every response
	r.Use(middleware.Recoverer)          // Recover from panics gracefully
	r.Use(mw.CacheControl(cfg.CacheTTL)) // HTTP cache headers
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   cfg.CORSOrigins,
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// --- Dependencies ---
	appCache := cache.New(cfg.CacheTTL)
	loader := data.NewLoader(cfg.DataDir)

	// --- Handlers ---
	healthH := handler.NewHealthHandler()
	buildingsH := handler.NewBuildingsHandler(loader, appCache)
	troopsH := handler.NewTroopsHandler(loader, appCache)
	faviconH := handler.NewFaviconHandler("static/favicon.ico")

	// --- Custom Error Handlers ---
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		handler.NotFound(w, "endpoint not found")
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		handler.Error(w, http.StatusMethodNotAllowed, "method not allowed")
	})

	// --- Routes ---
	r.Get("/", handler.RootHandler())
	r.Method("GET", "/health", healthH)
	r.Method("GET", "/favicon.ico", faviconH)

	// API routes
	r.Route("/api", func(r chi.Router) {
		// Building endpoints
		r.Get("/buildings", buildingsH.ListCategories)
		r.Get("/buildings/{category}", buildingsH.ListByCategory)
		r.Get("/buildings/{category}/{name}", buildingsH.GetBuilding)

		// Troop endpoints
		r.Get("/troops", troopsH.ListCategories)
		r.Get("/troops/{category}", troopsH.ListByCategory)
		r.Get("/troops/{category}/{name}", troopsH.GetTroop)
	})

	return r
}
