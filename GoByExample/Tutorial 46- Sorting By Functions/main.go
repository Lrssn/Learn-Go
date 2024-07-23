package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	lenCmp := func(a, b string) int { // Create a comparison function that takes two strings
		return cmp.Compare(len(a), len(b)) //Sort strings by length instead of alphabetically
	}

	slices.SortFunc(fruits, lenCmp) // SortFunc sorts array by the custom comparison function
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	slices.SortFunc(people,
		func(a, b Person) int { // Write function directly as a parameter to Sortfunc
			return cmp.Compare(a.age, b.age) // Sort structs by one value
		})
	fmt.Println(people)
}
