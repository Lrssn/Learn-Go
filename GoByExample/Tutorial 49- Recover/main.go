package main

import "fmt"

func mayPanic() {
	panic("a problem") // This function panics
}

func main() {

	defer func() {
		if r := recover(); r != nil { // Recover has to be defered so it runs after panic

			fmt.Println("Recovered. Error:\n", r) // Return value of recover is the error code from the panic
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()") // This code will not run due to the panic.
}
