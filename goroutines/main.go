package main

import (
	"fmt"
	"time"
)

func f1(a string) {
	for i := 0; i < 3; i++ {
		fmt.Println(a, ":", i)
	}
}

func main() {

	var f func(string)
	f = f1

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(1 * time.Millisecond)
	fmt.Println("done")
}
