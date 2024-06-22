package main

import "fmt"

func ping(pings chan<- string, msg string) { // Pings channel only sends values
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // Pongs channel can both send and recieve
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
