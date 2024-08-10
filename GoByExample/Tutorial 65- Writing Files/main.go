package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("tmp/dat1.txt", d1, 0644) // Write directly to file
	check(err)

	f, err := os.Create("tmp/dat2.txt") // Create a file and open
	check(err)

	defer f.Close() // Close file at end of function

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) // Write bytes to file
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n") // Write String to file
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f) // Create buffered write
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}
