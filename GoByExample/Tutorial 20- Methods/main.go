package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int { // func(name *object type) func-name(params) return-type
	return r.width * r.height // r is used as "self" in python i think
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
