package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K { // Generics are generic functions that take any type of input.
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct { // Linked lists can take any type
	head, tail *element[T]
}

type element[T any] struct { // Element can be any type
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) { // Push to end of list
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T { // get beginning of list and return a list of all items
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}
