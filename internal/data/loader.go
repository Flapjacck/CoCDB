// Package data handles loading and parsing JSON data files from the filesystem.
// It provides methods to list categories, enumerate items, and retrieve individual
// data entries while filtering out template files.
package data

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Loader reads and serves JSON data from a base directory.
type Loader struct {
	baseDir string
}

// NewLoader creates a Loader rooted at the given base directory.
func NewLoader(baseDir string) *Loader {
	return &Loader{baseDir: baseDir}
}

// CategoryInfo contains metadata about a data category (subdirectory).
type CategoryInfo struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Path  string `json:"path"`
}

// ItemSummary holds minimal identifying information for a data item.
type ItemSummary struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// ListCategories returns all subdirectories within the given sub-path,
// along with the count of JSON data files in each.
func (l *Loader) ListCategories(subPath string) ([]CategoryInfo, error) {
	dir := filepath.Join(l.baseDir, filepath.FromSlash(subPath))

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("directory not found: %s", subPath)
	}

	var categories []CategoryInfo
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		childPath := subPath + "/" + e.Name()
		files, _ := l.listJSONFiles(childPath)
		categories = append(categories, CategoryInfo{
			Name:  e.Name(),
			Count: len(files),
			Path:  childPath,
		})
	}
	return categories, nil
}

// ListItems returns all non-template JSON files within a category directory.
func (l *Loader) ListItems(subPath string) ([]ItemSummary, error) {
	files, err := l.listJSONFiles(subPath)
	if err != nil {
		return nil, err
	}

	var items []ItemSummary
	for _, f := range files {
		name := strings.TrimSuffix(f, ".json")
		items = append(items, ItemSummary{
			Name: name,
			Path: subPath + "/" + name,
		})
	}
	return items, nil
}

// GetItem reads and validates a single JSON file at the given sub-path.
// The sub-path should NOT include the .json extension.
func (l *Loader) GetItem(subPath string) (json.RawMessage, error) {
	target := filepath.Join(l.baseDir, filepath.FromSlash(subPath)+".json")

	// Prevent directory traversal by verifying resolved path stays within baseDir.
	absBase, err := filepath.Abs(l.baseDir)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve base directory")
	}
	absTarget, err := filepath.Abs(target)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve target path")
	}
	if !strings.HasPrefix(absTarget, absBase+string(filepath.Separator)) {
		return nil, fmt.Errorf("invalid path: %s", subPath)
	}

	raw, err := os.ReadFile(target)
	if err != nil {
		return nil, fmt.Errorf("item not found: %s", subPath)
	}

	if !json.Valid(raw) {
		return nil, fmt.Errorf("invalid JSON in: %s", subPath)
	}

	return json.RawMessage(raw), nil
}

// listJSONFiles returns filenames of non-template .json files in a directory.
func (l *Loader) listJSONFiles(subPath string) ([]string, error) {
	dir := filepath.Join(l.baseDir, filepath.FromSlash(subPath))

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("directory not found: %s", subPath)
	}

	var files []string
	for _, e := range entries {
		name := e.Name()
		if e.IsDir() || !strings.HasSuffix(name, ".json") || name == "template.json" {
			continue
		}
		files = append(files, name)
	}
	return files, nil
}
