package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1") // Create a environment variable
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	for _, e := range os.Environ() { // Print all env-vars
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
