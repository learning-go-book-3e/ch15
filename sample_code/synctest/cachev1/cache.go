package cachev1

import (
	"sync"
	"time"
)

// Cache is a simple in-memory key-value store where entries expire after a TTL.
type Cache[K comparable, V any] struct {
	mu    sync.Mutex
	items map[K]entry[V]
	ttl   time.Duration
}

type entry[V any] struct {
	value     V
	expiresAt time.Time
}

// New creates a Cache with the given TTL.
func New[K comparable, V any](ttl time.Duration) *Cache[K, V] {
	c := &Cache[K, V]{
		items: make(map[K]entry[V]),
		ttl:   ttl,
	}
	return c
}

// Set stores a value, resetting its expiration timer.
func (c *Cache[K, V]) Set(key K, val V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = entry[V]{
		value:     val,
		expiresAt: time.Now().Add(c.ttl),
	}
}

// Get retrieves a value. It returns the zero value and false if the key is
// missing or expired. If a key is expired, the value is deleted from the cache before
// returning the zero value and false.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	if !ok || time.Now().After(e.expiresAt) {
		delete(c.items, key)
		var zero V
		return zero, false
	}
	return e.value, true
}
