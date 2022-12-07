package scheduler

import (
	"sync"

	"github.com/leomirandadev/advance-golang-theory/fibonacci"
)

// How fast it will be?
// 10.000 * T / 16
func Task2() {

	// 16 CPUs
	var wg sync.WaitGroup
	wg.Add(TIMES)

	for i := 0; i < TIMES; i++ {
		go func() {
			execCalc()
			wg.Done()
		}()
	}

	wg.Wait()
}

func execCalc() {
	fibonacci.RecursionWithoutMemo(20)
	// 44418ns
	// 52531442ns
}
