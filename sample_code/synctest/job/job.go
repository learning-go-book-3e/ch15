package job

import "errors"

type Runner struct {
	slots chan struct{}
}

func New(max int) *Runner {
	return &Runner{
		slots: make(chan struct{}, max),
	}
}

var ErrMaxJobsReached = errors.New("maximum number of jobs reached")

func (r *Runner) Launch(f func()) error {
	select {
	case r.slots <- struct{}{}:
		go func() {
			defer func() { <-r.slots }()
			f()
		}()
		return nil
	default:
		return ErrMaxJobsReached
	}
}
