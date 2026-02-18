package handler

import (
	"log/slog"
	"net/http"

	"github.com/flapjacck/CoCDB/internal/cache"
	"github.com/flapjacck/CoCDB/internal/data"
	"github.com/go-chi/chi/v5"
)

// troopsBase is the data directory prefix for all troop data.
const troopsBase = "home_village/troops"

// TroopsHandler serves troop-related API endpoints.
type TroopsHandler struct {
	loader *data.Loader
	cache  *cache.Cache
}

// NewTroopsHandler creates a handler with the given data loader and cache.
func NewTroopsHandler(loader *data.Loader, c *cache.Cache) *TroopsHandler {
	return &TroopsHandler{loader: loader, cache: c}
}

// ListCategories handles GET /api/troops
// Returns all troop categories (elixir, dark_elixir, super) with item counts.
func (h *TroopsHandler) ListCategories(w http.ResponseWriter, r *http.Request) {
	const cacheKey = "troops:categories"

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	categories, err := h.loader.ListCategories(troopsBase)
	if err != nil {
		slog.Error("failed to list troop categories", "error", err)
		InternalError(w, "failed to load troop categories")
		return
	}

	h.cache.Set(cacheKey, categories)
	Success(w, categories, nil)
}

// ListByCategory handles GET /api/troops/{category}
// Returns all troops within a specific category.
func (h *TroopsHandler) ListByCategory(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	cacheKey := "troops:list:" + category

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	items, err := h.loader.ListItems(troopsBase + "/" + category)
	if err != nil {
		NotFound(w, "troop category not found: "+category)
		return
	}

	h.cache.Set(cacheKey, items)
	Success(w, items, nil)
}

// GetTroop handles GET /api/troops/{category}/{name}
// Returns full data for a specific troop.
func (h *TroopsHandler) GetTroop(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	name := chi.URLParam(r, "name")
	cacheKey := "troops:item:" + category + ":" + name

	if cached, ok := h.cache.Get(cacheKey); ok {
		Success(w, cached, nil)
		return
	}

	item, err := h.loader.GetItem(troopsBase + "/" + category + "/" + name)
	if err != nil {
		NotFound(w, "troop not found: "+name)
		return
	}

	h.cache.Set(cacheKey, item)
	Success(w, item, nil)
}
