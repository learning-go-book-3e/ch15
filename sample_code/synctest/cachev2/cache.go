package cachev2

import (
	"sync"
	"time"
)

type Stats struct {
	RemovedBySweep int
	RemovedByGet   int
}

// Cache is a simple in-memory key-value store where entries expire after a TTL.
type Cache[K comparable, V any] struct {
	mu    sync.Mutex
	items map[K]entry[V]
	ttl   time.Duration
	Stats Stats
	done  chan struct{}
}

type entry[V any] struct {
	value     V
	expiresAt time.Time
}

// New creates a Cache with the given TTL. A background goroutine reaps expired
// entries every sweepInterval.
func New[K comparable, V any](ttl, sweepInterval time.Duration) *Cache[K, V] {
	c := &Cache[K, V]{
		items: make(map[K]entry[V]),
		ttl:   ttl,
		done:  make(chan struct{}),
	}
	go c.clean(sweepInterval)
	return c
}

func (c *Cache[K, V]) Done() {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case <-c.done:
		return
	default:
		close(c.done)
	}
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
// missing or expired.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.items[key]
	if !ok || time.Now().After(e.expiresAt) {
		delete(c.items, key)
		if ok {
			c.Stats.RemovedByGet++
		}
		var zero V
		return zero, false
	}
	return e.value, true
}

// Len returns the number of items currently stored (including any that are
// expired but not yet reaped).
func (c *Cache[K, V]) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

// clean periodically removes expired entries.
//func (c *Cache[K, V]) clean(interval time.Duration) {
//	ticker := time.NewTicker(interval)
//	defer ticker.Stop()
//	for range ticker.C {
//		c.mu.Lock()
//		now := time.Now()
//		for k, e := range c.items {
//			if now.After(e.expiresAt) {
//				delete(c.items, k)
//				c.Stats.RemovedBySweep++
//			}
//		}
//		c.mu.Unlock()
//	}
//}

func (c *Cache[K, V]) clean(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for k, e := range c.items {
				if !now.Before(e.expiresAt) {
					delete(c.items, k)
					c.Stats.RemovedBySweep++
				}
			}
			c.mu.Unlock()
		case <-c.done:
			return
		}
	}
}
