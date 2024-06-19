package main

import "fmt"

func intSeq() func() int { // returns a function that returns a int
	i := 0
	return func() int { // return value is a function
		i++
		return i
	}
}

func main() {

	nextInt := intSeq() // nextInt is a function that returns an int

	fmt.Println(nextInt()) // the value of i is incremented each call
	fmt.Println(nextInt()) // the value of i in nextint is the same number getting
	fmt.Println(nextInt()) // incremented each call

	newInts := intSeq()
	fmt.Println(newInts()) // this increments a new i that is contained in newInts
}
