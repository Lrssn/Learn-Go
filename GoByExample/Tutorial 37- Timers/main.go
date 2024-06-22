package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second) // Create Timer

	<-timer1.C // Blocks until timer fires
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C // timer blocks thread
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop() // Cancel timer before it fires
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
