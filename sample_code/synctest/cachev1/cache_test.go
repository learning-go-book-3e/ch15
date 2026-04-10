package cachev1_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/learning-go-book-3e/ch15/sample_code/synctest/cachev1"
)

func TestCache_GetBeforeExpiry_standard(t *testing.T) {
	c := cachev1.New[string, string](1 * time.Second)
	c.Set("greeting", "hello")

	// Advance the clock to just before expiry.
	time.Sleep(990 * time.Millisecond)

	val, ok := c.Get("greeting")
	if !ok {
		t.Fatal("expected key to still be present before TTL")
	}
	if val != "hello" {
		t.Fatalf("got %q, want %q", val, "hello")
	}
	// Advance the clock to just after expiry.
	time.Sleep(20 * time.Millisecond)
	val, ok = c.Get("greeting")
	if ok {
		t.Fatal("expected key to not be present after TTL")
	}
}

func TestCache_GetAfterExpiry_standard(t *testing.T) {
	c := cachev1.New[string, int](1 * time.Second)
	c.Set("answer", 42)

	// Jump past the TTL.
	time.Sleep(1100 * time.Millisecond)

	_, ok := c.Get("answer")
	if ok {
		t.Fatal("expected key to be expired after TTL")
	}
}

func TestCache_SetResetsExpiry_standard(t *testing.T) {
	c := cachev1.New[string, string](1 * time.Second)
	c.Set("key", "v1")

	// Advance .99 seconds, then re-set the key to reset its TTL.
	time.Sleep(990 * time.Millisecond)
	c.Set("key", "v2")

	// Another .99 seconds — 1.98 total, but only .99 since the last Set.
	time.Sleep(990 * time.Millisecond)

	val, ok := c.Get("key")
	if !ok {
		t.Fatal("expected key to still be present after re-set")
	}
	if val != "v2" {
		t.Fatalf("got %q, want %q", val, "v2")
	}
}

func TestCache_GetBeforeExpiry_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev1.New[string, string](1 * time.Second)
		c.Set("greeting", "hello")

		// Advance the fake clock to just before expiry.
		time.Sleep(990 * time.Millisecond)

		val, ok := c.Get("greeting")
		if !ok {
			t.Fatal("expected key to still be present before TTL")
		}
		if val != "hello" {
			t.Fatalf("got %q, want %q", val, "hello")
		}
		// Advance the clock to just after expiry.
		time.Sleep(20 * time.Millisecond)
		val, ok = c.Get("greeting")
		if ok {
			t.Fatal("expected key to not be present after TTL")
		}
	})
}

func TestCache_GetAfterExpiry_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev1.New[string, int](1 * time.Second)
		c.Set("answer", 42)

		// Jump past the TTL.
		time.Sleep(1100 * time.Millisecond)

		_, ok := c.Get("answer")
		if ok {
			t.Fatal("expected key to be expired after TTL")
		}
	})
}

func TestCache_SetResetsExpiry_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		c := cachev1.New[string, string](1 * time.Second)
		c.Set("key", "v1")

		// Advance .99 seconds, then re-set the key to reset its TTL.
		time.Sleep(990 * time.Millisecond)
		synctest.Wait()
		c.Set("key", "v2")

		// Another .99 seconds — 1.98 total, but only .99 since the last Set.
		time.Sleep(990 * time.Millisecond)

		val, ok := c.Get("key")
		if !ok {
			t.Fatal("expected key to still be present after re-set")
		}
		if val != "v2" {
			t.Fatalf("got %q, want %q", val, "v2")
		}
	})
}
