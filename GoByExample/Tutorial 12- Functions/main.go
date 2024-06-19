package main

import "fmt"

func plus(a int, b int) int { // declaration name(values) return-value

	return a + b
}

func plusPlus(a, b, c int) int { // a,b,c are the same type
	return a + b + c
}

func main() {

	res := plus(1, 2) // initialize value with return
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
