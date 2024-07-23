package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops atomic.Uint64 // Atomic Unit64

	var wg sync.WaitGroup // Create a waitgroup

	for i := 0; i < 50; i++ {
		wg.Add(1) // Start 50 goroutines

		go func() {
			for c := 0; c < 1000; c++ {

				ops.Add(1) // Add 1 a thousand times on each routine
			}

			wg.Done() // Finish goroutine
		}()
	}

	wg.Wait() // Wait until all are finished

	fmt.Println("ops:", ops.Load()) // We get 50000 adds
}
