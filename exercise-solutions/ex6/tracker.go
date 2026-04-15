package ex6

import "sync"

// Tracker tracks the names of the users who access the resource.
type Tracker struct {
	mu    sync.Mutex
	names []string
}

// New creates a Tracker ready for use.
func New() *Tracker {
	return &Tracker{}
}

// Track adds the name of each accessing user to the list of names.
func (t *Tracker) Track(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.names = append(t.names, name)
}

// GetCount returns the current number of users who have accessed the resource.
func (t *Tracker) GetCount() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return len(t.names)
}

// AddIfLessThan is needed because you need to acquire a lock across both function calls,
// not individually. If you acquire the lock in GetCount and release it and then acquire
// the lock in Track, it's possible for two different goroutines to see GetCount return a
// value less than 200, both write succeed, and you end up with 201 or 202 instead of 200
// for the total number of writes.
func (t *Tracker) AddIfLessThan(max int, name string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if len(t.names) < max {
		t.names = append(t.names, name)
	}
}
