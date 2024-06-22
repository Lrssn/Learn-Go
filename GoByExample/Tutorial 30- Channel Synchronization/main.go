package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true // Sends a value to say thread is done
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done // Wait for message to finish app
}
