package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue { // range over the elements in a channel
		fmt.Println(elem)
	}
}
