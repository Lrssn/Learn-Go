package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	h := sha256.New() // Create a hash

	h.Write([]byte(s)) // Write expects bytes so strings need to be converted

	bs := h.Sum(nil) // This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing byte slice: it usually isn’t needed

	fmt.Println(s)
	fmt.Printf("%x\n", bs) // Print hash
}
