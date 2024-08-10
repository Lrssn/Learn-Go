package main

import (
	"embed"
)

// embed directives accept paths relative to the directory containing the Go source file
//
//go:embed tmp/single_file.txt
var fileString string // This directive embeds the contents of the file into the string variable immediately following it.

//go:embed tmp/single_file.txt
var fileByte []byte // Or embed the contents of the file into a []byte.

// We can also embed multiple files or even folders with wildcards.
//
//go:embed tmp/single_file.txt
//go:embed tmp/*.hash
var folder embed.FS // This uses a variable of the embed.FS type, which implements a simple virtual file system.

func main() {

	print(fileString) // Print out the contents of single_file.txt.
	print(string(fileByte))

	content1, _ := folder.ReadFile("tmp/file1.hash") // Retrieve some files from the embedded folder.
	print(string(content1))

	content2, _ := folder.ReadFile("tmp/file2.hash")
	print(string(content2))
}
