package main

import (
	"fmt"
	"maps" // add lib maps
)

func main() {

	m := make(map[string]int) // create map with map[key]value

	m["k1"] = 7 // key=k1 value=7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"] // get value from map using key=k1
	fmt.Println("v1:", v1)

	v3 := m["k3"] // if value dont exist returns zero value for type
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m)) // length is the number of entries in the map

	delete(m, "k2") // delete entry from map
	fmt.Println("map:", m)

	clear(m) // empty map
	fmt.Println("map:", m)

	_, prs := m["k2"] // second return value is bool if exists in map.
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} // initialize map with values
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) { // compare two maps
		fmt.Println("n == n2")
	}
}
