package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64) // Parse float64 from string
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) // Parse Int from string
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) // Parse Int from hex-formatted string
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64) // Parse Unsigned Int from string
	fmt.Println(u)

	k, _ := strconv.Atoi("135") // Parse Int from string
	fmt.Println(k)

	_, e := strconv.Atoi("wat") // returns error on fail
	fmt.Println(e)
}
