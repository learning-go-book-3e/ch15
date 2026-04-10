package job_test

import (
	"errors"
	"sync"
	"testing"
	"testing/synctest"

	"github.com/learning-go-book-3e/ch15/sample_code/synctest/job"
)

func TestJobRunner(t *testing.T) {
	out := make(chan int, 5)
	wait := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(5)
	jr := job.New(5)
	// put in 5 jobs without error
	for i := range 5 {
		err := jr.Launch(func() {
			<-wait
			out <- i
			wg.Done()
		})
		if err != nil {
			t.Error("unexpected error", err)
		}
	}
	// put in 6th job
	err := jr.Launch(func() {
		t.Log("another job, won't be scheduled")
	})
	if !errors.Is(err, job.ErrMaxJobsReached) {
		t.Error("unexpected error", err)
	}
	close(wait)
	wg.Wait()
	close(out)
	total := 0
	for i := range out {
		total += i
	}
	// make sure it's the sum of 0 to 4 (10)
	if total != 10 {
		t.Error("missed a job", total)
	}
}

func TestJobRunner_synctest(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		out := make([]bool, 5)
		jr := job.New(5)
		// put in 5 jobs without error
		for i := range 5 {
			err := jr.Launch(func() {
				out[i] = true
				t.Log(i, "done")
			})
			if err != nil {
				t.Error("unexpected error", err)
			}
		}
		t.Log("launching error job")
		// put in 6th job
		err := jr.Launch(func() {
			t.Log("this job should not run")
		})
		t.Log("done launching error job")
		if !errors.Is(err, job.ErrMaxJobsReached) {
			t.Error("unexpected error", err)
		}
		t.Log("waiting")
		synctest.Wait()
		t.Log("finished wait")
		// confirm all jobs completed
		for i := range out {
			if !out[i] {
				t.Error("should have completed job", i)
			}
		}
	})
}
