package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Print(rand.IntN(100), ",") // rand.IntN(x) returns [0-100[
	fmt.Print(rand.IntN(100))
	fmt.Println()

	fmt.Println(rand.Float64()) // Random [0.0f-1.0f[

	fmt.Print((rand.Float64()*5)+5, ",") // [5.0f-10.0f[
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	s2 := rand.NewPCG(42, 1024)  // Create a known seed
	r2 := rand.New(s2)           // Create random engine using seed
	fmt.Print(r2.IntN(100), ",") // Get random number using engine
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024) // Using the same seed nets the same results
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
