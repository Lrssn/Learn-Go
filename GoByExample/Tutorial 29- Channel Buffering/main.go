package main

import "fmt"

func main() {

	messages := make(chan string, 2) // make a channel with 2 values

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages) // Buffered messages are first in first out
	fmt.Println(<-messages)
}
