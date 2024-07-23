package main

import "os"

func main() {

	panic("a problem") // Panic crashes the app with a message

	_, err := os.Create("/tmp/file") // If a function returns a error is a usual usecase for panic
	if err != nil {
		panic(err) // Panic with the error as a message
	}
}
