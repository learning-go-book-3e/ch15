package ex6

import (
	"sync"
	"testing"
)

func TestTracker(t *testing.T) {
	tracker := New()

	var wg sync.WaitGroup
	for _, s := range []string{"Fred", "Mary", "Pat"} {
		wg.Go(func() {
			for range 100 {
				if tracker.GetCount() < 200 {
					tracker.Track(s)
				}
			}
		})
	}

	wg.Wait()
	if count := tracker.GetCount(); count != 200 {
		t.Errorf("expected 200, got %d", count)
	}
}
