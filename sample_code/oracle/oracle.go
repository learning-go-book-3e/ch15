package oracle

import (
	"context"
	"log/slog"
)

type Request struct {
	Query    string
	Response chan string
}

func Launch(ctx context.Context) chan<- Request {
	ch := make(chan Request)
	go func() {
		for {
			select {
			case request := <-ch:
				// a very agreeable oracle
				request.Response <- request.Query + " Yes!"
			case <-ctx.Done():
				slog.Info("context canceled")
				return
			}
		}
	}()
	return ch
}
