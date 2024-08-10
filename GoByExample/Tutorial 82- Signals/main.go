package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 1) // Channel to recieve os signals

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // The specific signals to look for

	done := make(chan bool, 1)

	go func() {

		sig := <-sigs // Goroutine that waits for signal
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
