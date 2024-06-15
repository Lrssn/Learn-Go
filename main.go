package main

import "fmt"

func vals() (int, int) { // keyword for function, name(values) (return values)
	return 3, 7
}

func main() {

	a, b := vals() // initialize both values from function return values
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // skip the first value and only use second
	fmt.Println(c)
}
