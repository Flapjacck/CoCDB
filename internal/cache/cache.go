// Package cache provides a thread-safe in-memory key-value cache with TTL support.
// Expired entries are automatically evicted by a background goroutine.
package cache

import (
	"sync"
	"time"
)

// entry represents a single cached item with an expiration timestamp.
type entry struct {
	value     interface{}
	expiresAt time.Time
}

// Cache is a concurrent-safe in-memory cache with automatic TTL-based expiration.
type Cache struct {
	mu      sync.RWMutex
	entries map[string]entry
	ttl     time.Duration
	stop    chan struct{}
}

// New creates a Cache with the given TTL and starts a background eviction loop.
func New(ttl time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]entry),
		ttl:     ttl,
		stop:    make(chan struct{}),
	}
	go c.evictLoop()
	return c
}

// Get retrieves a value by key. Returns (value, true) if found and not expired,
// or (nil, false) otherwise.
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	e, exists := c.entries[key]
	if !exists || time.Now().After(e.expiresAt) {
		return nil, false
	}
	return e.value, true
}

// Set stores a value in the cache with the configured TTL.
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = entry{
		value:     value,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Delete removes a single entry by key.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.entries, key)
}

// Flush removes all entries from the cache.
func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries = make(map[string]entry)
}

// Size returns the current number of entries in the cache.
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.entries)
}

// Close stops the background eviction goroutine.
func (c *Cache) Close() {
	close(c.stop)
}

// evictLoop periodically scans and removes expired entries.
// It runs at half the TTL interval for timely cleanup.
func (c *Cache) evictLoop() {
	ticker := time.NewTicker(c.ttl / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for key, e := range c.entries {
				if now.After(e.expiresAt) {
					delete(c.entries, key)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}
