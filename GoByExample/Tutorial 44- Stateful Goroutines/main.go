package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct { // Create structs for data operations
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	var readOps uint64 // Create counters
	var writeOps uint64

	reads := make(chan readOp) // Create channels
	writes := make(chan writeOp)

	go func() { // Create a goroutine that owns the data
		var state = make(map[int]int) // In this case a map
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key] // On ReadOp, return the value of the data
			case write := <-writes:
				state[write.key] = write.val // On WriteOp, change the value of the data
				write.resp <- true           // Return a response
			}
		}
	}()

	for r := 0; r < 100; r++ { // Create 100 ReadOps
		go func() {
			for {
				read := readOp{ // Create ReadOp
					key:  rand.Intn(5),   // Random value between 0 and 4
					resp: make(chan int)} // Channel that takes a int
				reads <- read                 // Queue ReadOp
				<-read.resp                   // Dequeue into nothing
				atomic.AddUint64(&readOps, 1) // Increment counter
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ { // Create 10 WriteOps
		go func() {
			for {
				write := writeOp{ // Create WriteOps
					key:  rand.Intn(5),    // Random Value between 0 and 4
					val:  rand.Intn(100),  // Random Value between 0 and 99
					resp: make(chan bool)} // Channel that takes a bool
				writes <- write                // Queue WriteOp
				<-write.resp                   // Dequeue response into nothing
				atomic.AddUint64(&writeOps, 1) // Increment counter
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps) // Read counters
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
