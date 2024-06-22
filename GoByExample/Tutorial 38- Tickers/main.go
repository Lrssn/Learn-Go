package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond) // tickers fire every x seconds
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done: // thread runs until done
				return
			case t := <-ticker.C: // prints every tick
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
