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
				tracker.AddIfLessThan(200, s)
			}
		})
	}

	wg.Wait()
	if count := tracker.GetCount(); count != 200 {
		t.Errorf("expected 200, got %d", count)
	}
}
