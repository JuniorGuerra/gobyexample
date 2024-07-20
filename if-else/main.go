package main

import "fmt"

func main() {
	if true {
		fmt.Println("true")
	}
	if !false {
		fmt.Println("!false")
	}
	if ok := true; ok {
		fmt.Println(ok)
	}
	if ok := false; ok {
		fmt.Println("is not ok")
	} else if !ok && ok {
		fmt.Println("ok is confused")
	} else {
		fmt.Println("ok")
	}

}
