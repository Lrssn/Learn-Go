package main

import "fmt"

func main() {

	var a = "initial" // string
	fmt.Println(a)

	var b, c int = 1, 2 // multiple ints
	fmt.Println(b, c)

	var d = true // bool
	fmt.Println(d)

	var e int // empty int
	fmt.Println(e)

	f := "apple" // auto infer type
	fmt.Println(f)
}
