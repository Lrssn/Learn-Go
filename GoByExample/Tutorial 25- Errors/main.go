package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {

		return -1, errors.New("can't work with 42") // this is how you create a error
	}

	return arg + 3, nil // errors == nil means there is no error
}

var ErrOutOfTea = fmt.Errorf("no more tea available") // you can create errors as varialbles
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {

		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {

		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {

			if errors.Is(err, ErrOutOfTea) { // errors.Is checks list of errors for a specific one
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
}
