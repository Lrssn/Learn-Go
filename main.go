package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1) // this function calls itself
}

func main() {
	fmt.Println(fact(7))

	var fib func(n int) int // closures can be recursive if you declare them before usage
	// like in c++ you have to declare functions before usage

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2) // function that calls its own clusure function
	}

	fmt.Println(fib(7))
}
