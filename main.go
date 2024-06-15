package main

import "fmt"

func sum(nums ...int) { // variable number of ints. not set at compile time. could be zero?
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2) // call the function with how many variables youd like
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...) // call function with all values in a slice
}
