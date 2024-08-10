package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	p := filepath.Join("dir1", "dir2", "filename") // Create a filepath
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))       // Always use join instead of slashes
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // Join normalizes filepaths

	fmt.Println("Dir(p):", filepath.Dir(p))   // Directory to file
	fmt.Println("Base(p):", filepath.Base(p)) // Filename

	fmt.Println(filepath.IsAbs("dir/file")) // Check if file is absolute
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	ext := filepath.Ext(filename) // Get file extension
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext)) // Remove extension to get filename

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
