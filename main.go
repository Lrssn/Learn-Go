package main

import "fmt"

func main() {

	if 8%2 == 0 || 7%2 == 0 { // ||=OR, &&=AND
		fmt.Println("either 8 or 7 are even")
	}
	if 8%2 == 0 && 6%2 == 0 { // ||=OR, &&=AND
		fmt.Println("both 8 and 6 are even")
	}

	if num := 9; num < 0 { // if-else if-else
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
