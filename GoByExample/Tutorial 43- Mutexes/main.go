package main

import (
	"fmt"
	"sync"
)

type Container struct { // Container holds a map of counters
	mu       sync.Mutex // Add a mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()         // Lock the mutex
	defer c.mu.Unlock() // Unlocks it at the end of the function, defer runs att the end of function
	c.counters[name]++  // Increase the counter
}

func main() {
	c := Container{ // Create and initialize container

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup // Create waitgroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name) // Increment the map
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000) // access the same counter
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
