package main

import (
	"fmt"
	"os"
)

func main() {

	argsWithProg := os.Args // os.Args hold the arguments
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3] // Get args with index

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
