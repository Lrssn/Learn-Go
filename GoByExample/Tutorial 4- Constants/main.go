package main

import (
	"fmt"
	"math"
)

const s string = "constant" // global constant

func main() {

	fmt.Println(s)
	const n = 500000000 // local constant

	const d = 3e20 / n // 3*10^20
	fmt.Println(d)

	fmt.Println(int64(d))    // d set to 64bit integer. prints without exponential
	fmt.Println(math.Sin(n)) // math has functions i guess
}
