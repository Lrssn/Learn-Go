package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct { // Custom types
	Page   int
	Fruits []string
}

type response2 struct { // Only exported fields will be encoded/decoded in JSON. Fields must start with capital letters to be exported.
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true) // Bool
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1) // Int
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34) // Float
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher") // String
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} // Slice
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7} // Map
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{ // Set values of type
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D) // Parse it to JSON
	fmt.Println(string(res1B))

	res2D := &response2{ // Set values of type
		Page:   1, // The names come from the type and is not the variable name
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D) // Parse it to JSON
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil { // Decode JSON to a map and check for errors
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // Get data from map to a variable as specific type
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // Accessing nested data needs a conversion
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res) // Decode into custom type
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) // Decode directly to os
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
