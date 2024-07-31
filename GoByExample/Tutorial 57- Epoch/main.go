package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now() // Time now
	fmt.Println(now)

	fmt.Println(now.Unix())      // Time since UNIX Starttime as a variable in seconds
	fmt.Println(now.UnixMilli()) // in Milliseconds
	fmt.Println(now.UnixNano())  // in Nanoseconds

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
