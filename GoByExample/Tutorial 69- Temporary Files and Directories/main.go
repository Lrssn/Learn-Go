package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.CreateTemp("", "sample") // Create temporary file with name sampleXXX
	check(err)                            // It adds a random number as XXX

	fmt.Println("Temp file name:", f.Name())

	defer os.Remove(f.Name()) // Remove file

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "sampledir") // Create temporary file with name sampleXXX
	check(err)                                  // It adds a random number as XXX
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname) // Remove all temporary dirs

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666) // Create file in temporary dir
	check(err)
}
