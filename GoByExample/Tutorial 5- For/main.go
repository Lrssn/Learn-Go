package main

import "fmt"

func main() {

	i := 1 // loop with existing value
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ { // loop with internal value
		fmt.Println(j)
	}

	for i := range 3 { // loop with range
		fmt.Println("range", i)
	}

	for { // loop will run forever/ until break
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue // goto next instance of loop
		}
		fmt.Println(n)
	}
}
