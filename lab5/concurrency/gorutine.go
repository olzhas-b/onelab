package concurrency

import (
	"errors"
	"sync"
)

func Execute(task []func() error, N int, E int) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var outOfLimits = errors.New("ERRORS")
	capacity := make(chan int, N)
	for i := 0; i < len(task); i++ {
		i := i
		mu.Lock()
		if E <= 0 {
			break
		}
		mu.Unlock()
		wg.Add(1)
		capacity <- 1
		go func() {
			defer wg.Done()
			if err := task[i](); err != nil {
				mu.Lock()
				E--
				mu.Unlock()
			}
			<- capacity
		}()
	}
	wg.Wait()
	if E <= 0 {
		return outOfLimits
	}
	return nil
}