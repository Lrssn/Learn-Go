package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) { // Helper for checking for errors
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.ReadFile("tmp/dat.txt") // Read whole file into memory
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("tmp/dat.txt") // Open file
	check(err)

	b1 := make([]byte, 5) // Read up to 5 bytes into memory
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, io.SeekStart) // Read from specific position from start of file
	check(err)
	b2 := make([]byte, 2) // Read 2 bytes
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(4, io.SeekCurrent) // Read from relative to cursor
	check(err)

	_, err = f.Seek(-10, io.SeekEnd) // Read from End
	check(err)

	o3, err := f.Seek(6, io.SeekStart) // Read from start
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	r4 := bufio.NewReader(f) // Buffered reader
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // Close file
}
