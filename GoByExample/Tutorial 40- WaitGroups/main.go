package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup // Create waitgroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment work group

		go func() {
			defer wg.Done() // Add goroutine to waitgroup
			worker(i)       // Run a function from goroutine
		}()
	}

	wg.Wait() // Block until all goroutine in waitgroup is finished

}
