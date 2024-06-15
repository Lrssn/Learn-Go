package main

import (
	"fmt"
	"slices" // add slices library
)

func main() {

	var s []string // slices are like c++ <vectors>
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3) // create a slice with 3 (uninitialized = nil/zero)values
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a" // set values like in arrays
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) // length is now a function and not a property

	s = append(s, "d") // add values to the end of slice
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // create a slice c with same length as s
	copy(c, s)                  // copy the values from s to c
	fmt.Println("cpy:", c)

	l := s[2:5] // create slice with the values 3-6 from slice s
	fmt.Println("sl1:", l)

	l = s[:5] // create slice with all values until 6 from slice s
	fmt.Println("sl2:", l)

	l = s[2:] // create slice with all values from 3 from slice s
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // declare slice t and initialize the values
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) { // compare two slices
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) // 2D-slices are weird and can contain different amounts of values, WEIRD
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
