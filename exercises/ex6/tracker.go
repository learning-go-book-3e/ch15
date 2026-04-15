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
	t.names = append(t.names, name)
}

// GetCount returns the current number of users who have accessed the resource.
func (t *Tracker) GetCount() int {
	return len(t.names)
}
