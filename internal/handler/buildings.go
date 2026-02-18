package handler

import (
	"log/slog"
	"net/http"

	"github.com/flapjacck/CoCDB/internal/cache"
	"github.com/flapjacck/CoCDB/internal/data"
	"github.com/go-chi/chi/v5"
)

// buildingsBase is the data directory prefix for all building data.
const buildingsBase = "home_village/buildings"

// BuildingsHandler serves building-related API endpoints.
type BuildingsHandler struct {
	loader *data.Loader
	cache  *cache.Cache
}

// NewBuildingsHandler creates a handler with the given data loader and cache.
func NewBuildingsHandler(loader *data.Loader, c *cache.Cache) *BuildingsHandler {
	return &BuildingsHandler{loader: loader, cache: c}
}

// ListCategories handles GET /api/buildings
// Returns all building categories (army, defensive, resource, traps) with item counts.
func (h *BuildingsHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	const cacheKey = "buildings:categories"

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	categories, err := h.loader.ListCategories(buildingsBase)
	if err != nil {
		slog.Error("failed to list building categories", "error", err)
		InternalError(w, "failed to load building categories")
		return
	}

	h.cache.Set(cacheKey, categories)
	Success(w, categories, nil)
}

// ListByCategory handles GET /api/buildings/{category}
// Returns all buildings within a specific category.
func (h *BuildingsHandler) ListByCategory(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	cacheKey := "buildings:list:" + category

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	items, err := h.loader.ListItems(buildingsBase + "/" + category)
	if err != nil {
		NotFound(w, "building category not found: "+category)
		return
	}

	h.cache.Set(cacheKey, items)
	Success(w, items, nil)
}

// GetBuilding handles GET /api/buildings/{category}/{name}
// Returns full data for a specific building.
func (h *BuildingsHandler) GetBuilding(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	name := chi.URLParam(r, "name")
	cacheKey := "buildings:item:" + category + ":" + name

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	item, err := h.loader.GetItem(buildingsBase + "/" + category + "/" + name)
	if err != nil {
		NotFound(w, "building not found: "+name)
		return
	}

	h.cache.Set(cacheKey, item)
	Success(w, item, nil)
}
