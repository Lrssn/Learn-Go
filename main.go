package main

import "fmt"

func main() {

	var a [5]int // array with 5 ints, ints defaults to 0
	fmt.Println("emp:", a)

	a[4] = 100 // set index 4(entry no 5 last entry) starts at 0
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // get array length

	b := [5]int{1, 2, 3, 4, 5} // initialize array with values
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5} // compiler sets the length of the array
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} // set [0]=100 [3]=400 [4]=500, [1] and [2] gets set to 0
	fmt.Println("idx:", b)

	var twoD [2][3]int // initialize 2D-array
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j // set values in 2D-array
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{ // initialize values in 2D-array
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
