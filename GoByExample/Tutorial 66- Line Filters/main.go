package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin) // Create buffered scanner

	for scanner.Scan() { // Read to scanner

		ucl := strings.ToUpper(scanner.Text()) // Add scanner text to string and make it uppercase

		fmt.Println(ucl)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
