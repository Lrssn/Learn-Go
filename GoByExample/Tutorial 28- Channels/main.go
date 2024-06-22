package main

import "fmt"

func main() {

	messages := make(chan string) // Create a new Channel into a variable

	go func() { messages <- "ping" }() // Use "<-"-syntax to send a variable into a channel

	msg := <-messages // use the same syntax to get a variable out of the channel
	fmt.Println(msg)
}
