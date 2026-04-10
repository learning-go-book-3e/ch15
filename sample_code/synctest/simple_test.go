package synctest

import (
	"testing"
	"testing/synctest"
	"time"
)

func TestTimeWaiting(t *testing.T) {
	start := time.Now() // real wall clock time/date
	go func() {
		time.Sleep(1 * time.Second)
		t.Log(time.Since(start)) // always logs "1s"
		t.Log(time.Now().UTC())
	}()
	time.Sleep(2 * time.Second) // the goroutine above will run before this Sleep returns
	t.Log(time.Since(start))    // always logs "2s"
	t.Log(time.Now().UTC())
}

func TestTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now() // always midnight UTC 2000-01-01
		go func() {
			time.Sleep(1 * time.Second)
			t.Log(time.Since(start)) // always logs "1s"
			t.Log(time.Now().UTC())
		}()
		time.Sleep(2 * time.Second) // the goroutine above will run before this Sleep returns
		t.Log(time.Since(start))    // always logs "2s"
		t.Log(time.Now().UTC())
	})
}
