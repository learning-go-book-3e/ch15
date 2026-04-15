package race

import "sync"

func getCounter() int {
	var counter int
	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() {
			for range 1000 {
				counter++
			}
		})
	}
	wg.Wait()
	return counter
}
