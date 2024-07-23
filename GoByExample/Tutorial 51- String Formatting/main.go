package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p) // Struct values

	fmt.Printf("struct2: %+v\n", p) // Struct with keys and values

	fmt.Printf("struct3: %#v\n", p) // Struct with variable name, keys and values

	fmt.Printf("type: %T\n", p) // Type

	fmt.Printf("bool: %t\n", true) // Bool value

	fmt.Printf("int: %d\n", 123) // Integer in base 10

	fmt.Printf("bin: %b\n", 14) // Integer in bas 2

	fmt.Printf("char: %c\n", 33) // Ascii value

	fmt.Printf("hex: %x\n", 456) // Hex value

	fmt.Printf("float1: %f\n", 78.9) // Float value

	fmt.Printf("float2: %e\n", 123400000.0) // Different Scientific notation
	fmt.Printf("float3: %E\n", 123400000.0)

	fmt.Printf("str1: %s\n", "\"string\"") // Basic string

	fmt.Printf("str2: %q\n", "\"string\"") // The full string

	fmt.Printf("str3: %x\n", "hex this") // Text as their Ascii value 2 numbers(XX)

	fmt.Printf("pointer: %p\n", &p) // The pointer adress

	fmt.Printf("width1: |%6d|%6d|\n", 12, 345) // Setting the width of a value, right-justified by default

	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45) // Same for floats

	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45) // To left-justify use "-"

	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b") // Same for strings

	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string") // Sprintf has the string as a returnvalue without printing to console
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
