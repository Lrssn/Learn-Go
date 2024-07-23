package main

import (
	"fmt"
	"slices" // Slices contains functions for sorting
)

func main() {

	strs := []string{"c", "a", "b"} // Create array of strings
	slices.Sort(strs)               // Sort array based on type, Strings alphabetically
	fmt.Println("Strings:", strs)   // The original order is forgotten

	ints := []int{7, 2, 4}
	slices.Sort(ints) // Sorts ints numerically
	fmt.Println("Ints:   ", ints)

	s := slices.IsSorted(ints) // Check if it is sorted
	fmt.Println("Sorted: ", s)
}
