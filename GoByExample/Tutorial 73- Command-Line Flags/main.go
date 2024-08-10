package main

import (
	"flag"
	"fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string") // Default flag declaration

	numbPtr := flag.Int("numb", 42, "an int")     // Default flag declaration
	forkPtr := flag.Bool("fork", false, "a bool") // Default flag declaration

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") // Connect a flag to a variable

	flag.Parse() // When all flags are declared you call parse

	fmt.Println("word:", *wordPtr) // Dereference the pointers with *x to get the actual values.
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) // Get the remaining flags
}
