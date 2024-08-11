package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 32 {
		return -1, errors.New("can't work with 32")
	}
	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPorwer = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPorwer)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 32} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("we should buy new tea")
			} else if errors.Is(err, ErrPorwer) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
	}
	fmt.Println("Tea is ready!!")
}
