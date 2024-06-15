package main

import "fmt"

func zeroval(ival int) { // copies ival
	ival = 0 // sets copy to 0 not changing original
}

func zeroptr(iptr *int) {
	*iptr = 0 // changes original
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i) // get address of i
}
