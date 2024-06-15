package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // range from slice = len(slice)
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // return values for arrays and slices are [index, value]
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { // return value for maps are [key, value], range is number of entries
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // with only one return value you get the keys from map
		fmt.Println("key:", k)
	}

	for i, c := range "go" { // range in strings returns [index, unicode value]
		fmt.Println(i, c)
	}
}
