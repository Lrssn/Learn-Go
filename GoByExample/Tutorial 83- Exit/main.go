package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("!") // Defered functions dont run after exit

	os.Exit(3) // Exit with specific status
}
