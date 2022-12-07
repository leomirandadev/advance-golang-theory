package scheduler

import (
	"sync"
	"time"
)

// How fast it will be?
// 1 second
func Task1() {

	// 16 CPUs
	var wg sync.WaitGroup
	wg.Add(TIMES)

	for i := 0; i < TIMES; i++ {
		go func() {
			execBlocked()
			wg.Done()
		}()
	}

	wg.Wait()
}

// sleep
// http requests
// io syscalls
func execBlocked() {
	time.Sleep(1 * time.Second) // 1s
}
