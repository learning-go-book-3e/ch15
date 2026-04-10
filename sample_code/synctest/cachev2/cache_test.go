package cachev2_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/learning-go-book-3e/ch15/sample_code/synctest/cachev2"
)

func TestCache_GetBeforeExpiry(t *testing.T) {
	// synctest.Test runs our function inside a "bubble" where time is
	// virtualised. time.Now, time.Sleep, time.NewTicker, etc. all use a
	// fake clock that advances instantly when every goroutine in the
	// bubble is blocked.
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](5*time.Second, 1*time.Second)
		defer c.Done()
		c.Set("greeting", "hello")

		// Advance the fake clock to just before expiry.
		time.Sleep(4 * time.Second)

		// Wait for all background goroutines (the reaper) to settle.
		synctest.Wait()

		val, ok := c.Get("greeting")
		if !ok {
			t.Fatal("expected key to still be present before TTL")
		}
		if val != "hello" {
			t.Fatalf("got %q, want %q", val, "hello")
		}
	})
}

func TestCache_GetAfterExpiry(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, int](5*time.Second, 1*time.Second)
		defer c.Done()
		c.Set("answer", 42)

		// Jump past the TTL.
		time.Sleep(6 * time.Second)
		synctest.Wait()

		_, ok := c.Get("answer")
		if ok {
			t.Fatal("expected key to be expired after TTL")
		}
	})
}

func TestCache_ReaperRemovesExpiredEntries(t *testing.T) {
	// The reaper goroutine fires on its own ticker. Without synctest
	// we'd need a real sleep (slow) or a fake clock interface (noisy).
	// Inside the bubble the ticker fires instantly.
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](2*time.Second, 1*time.Second)
		defer c.Done()

		c.Set("a", "alpha")
		c.Set("b", "bravo")

		if c.Len() != 2 {
			t.Fatalf("expected 2 items, got %d", c.Len())
		}

		// Advance past the TTL so both entries expire, then let the
		// reaper tick at least once.
		time.Sleep(3 * time.Second)
		synctest.Wait()

		if n := c.Len(); n != 0 {
			t.Fatalf("expected reaper to remove all entries, but %d remain", n)
		}
	})
}

func TestCache_SetResetsExpiry(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](5*time.Second, 1*time.Second)
		defer c.Done()
		c.Set("key", "v1")

		// Advance 4 seconds, then re-set the key to reset its TTL.
		time.Sleep(4 * time.Second)
		synctest.Wait()
		c.Set("key", "v2")

		// Another 4 seconds — 8s total, but only 4s since the last Set.
		time.Sleep(4 * time.Second)
		synctest.Wait()

		val, ok := c.Get("key")
		if !ok {
			t.Fatal("expected key to still be present after re-set")
		}
		if val != "v2" {
			t.Fatalf("got %q, want %q", val, "v2")
		}
	})
}

func TestCache_Clean(t *testing.T) {
	c := cachev2.New[string, string](2*time.Second, 1*time.Second)
	c.Set("key", "hello")
	val, ok := c.Get("key")
	if !ok {
		t.Fatal("expected key to be present")
	}
	if val != "hello" {
		t.Fatalf("got %q, want %q", val, "hello")
	}
	time.Sleep(4 * time.Second)
	val, ok = c.Get("key")
	if ok {
		t.Fatal("expected key to not be present")
	}
	if c.Stats.RemovedByGet != 0 {
		t.Error("expected removed by Get to be 0")
	}
	if c.Stats.RemovedBySweep != 1 {
		t.Error("expected removed by Sweep to be 1")
	}
}

func TestCache_Goroutines(t *testing.T) {
	c := cachev2.New[string, string](2*time.Second, 1*time.Second)
	t.Cleanup(c.Done)
	go func() {
		time.Sleep(1 * time.Second)
		c.Set("key", "hello")
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c.Set("key2", "world")
	}()
	time.Sleep(1 * time.Second)
	t.Log("first wait")
	if v, ok := c.Get("key"); v != "hello" || !ok {
		t.Errorf("expected hello and true for key, got %s and %v", v, ok)
	}
	if v, ok := c.Get("key2"); v != "" || ok {
		t.Errorf("expected blank and false for key2, got %s and %v", v, ok)
	}
	time.Sleep(1 * time.Second)
	t.Log("second wait")
	if v, ok := c.Get("key"); v != "hello" || !ok {
		t.Errorf("expected hello and true for key, got %s and %v", v, ok)
	}
	if v, ok := c.Get("key2"); v != "world" || !ok {
		t.Errorf("expected world and true for key2, got %s and %v", v, ok)
	}
	time.Sleep(1 * time.Second)
	t.Log("third wait")
	if v, ok := c.Get("key"); v != "" || ok {
		t.Errorf("expected blank and false for key, got %s and %v", v, ok)
	}
	if v, ok := c.Get("key2"); v != "world" || !ok {
		t.Errorf("expected world and true for key2, got %s and %v", v, ok)
	}
}

func TestCache_Goroutines_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev2.New[string, string](2*time.Second, 1*time.Second)
		t.Cleanup(c.Done)
		go func() {
			time.Sleep(1 * time.Second)
			c.Set("key", "hello")
		}()
		go func() {
			time.Sleep(2 * time.Second)
			c.Set("key2", "world")
		}()
		if time.Now().UTC().String() != "2000-01-01 00:00:00 +0000 UTC" {
			t.Error("expected 2000-01-01 00:00:00 +0000 UTC for initial time")
		}
		time.Sleep(1 * time.Second)
		synctest.Wait()
		t.Log("first wait")
		if time.Now().UTC().String() != "2000-01-01 00:00:01 +0000 UTC" {
			t.Error("expected 2000-01-01 00:00:00 +0000 UTC for first wait")
		}
		if v, ok := c.Get("key"); v != "hello" || !ok {
			t.Errorf("expected hello and true for key, got %s and %v", v, ok)
		}
		if v, ok := c.Get("key2"); v != "" || ok {
			t.Errorf("expected blank and false for key2, got %s and %v", v, ok)
		}
		time.Sleep(1 * time.Second)
		synctest.Wait()
		t.Log("second wait")
		if time.Now().UTC().String() != "2000-01-01 00:00:02 +0000 UTC" {
			t.Error("expected 2000-01-01 00:02:00 +0000 UTC for second wait")
		}
		if v, ok := c.Get("key"); v != "hello" || !ok {
			t.Errorf("expected hello and true for key, got %s and %v", v, ok)
		}
		if v, ok := c.Get("key2"); v != "world" || !ok {
			t.Errorf("expected world and true for key2, got %s and %v", v, ok)
		}
		time.Sleep(1 * time.Second)
		synctest.Wait()
		t.Log("third wait")
		if time.Now().UTC().String() != "2000-01-01 00:00:03 +0000 UTC" {
			t.Error("expected 2000-01-01 00:03:00 +0000 UTC for third wait")
		}
		if v, ok := c.Get("key"); v != "" || ok {
			t.Errorf("expected blank and false for key, got %s and %v", v, ok)
		}
		if v, ok := c.Get("key2"); v != "world" || !ok {
			t.Errorf("expected world and true for key2, got %s and %v", v, ok)
		}
	})
}
