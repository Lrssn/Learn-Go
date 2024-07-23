package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") // Direct Check if it matches
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") // Optimized Regex

	fmt.Println(r.MatchString("peach")) // Functions are available on struct

	fmt.Println(r.FindString("peach punch")) // The first part that matches the Regex

	fmt.Println("idx:", r.FindStringIndex("peach punch")) // Find the first index where it matches

	fmt.Println(r.FindStringSubmatch("peach punch")) // Find all parts that matches all parts in regex

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // Find indexes for all parts

	fmt.Println(r.FindAllString("peach punch pinch", -1)) // Find all matches

	fmt.Println("all:", r.FindAllStringSubmatchIndex( // Find all indexes
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2)) // Find first 2 matches, second value defaults to 1, -1 gives all

	fmt.Println(r.Match([]byte("peach"))) // We can also find bytewise matches

	r = regexp.MustCompile("p([a-z]+)ch") // Panics if it encounters an error
	fmt.Println("regexp:", r)

	fmt.Println(r.ReplaceAllString("a peach and a peach", "<fruit>")) // Replaces all matches with 2 value

	in := []byte("a peach and a peach.")
	out := r.ReplaceAllFunc(in, bytes.ToUpper) // Make matches uppercase
	fmt.Println(string(out))
}
