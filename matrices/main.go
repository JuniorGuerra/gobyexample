package main

import (
	"fmt"
	"math"
)

func staticMa() {
	var ma [9][3]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 3; j++ {
			ma[i][j] = int(math.Pow(float64(i+1), float64(j+1)))
		}
	}
	fmt.Println(ma)
}

func multidimensionalArray() {
	staticMa()
}

func main() {
	multidimensionalArray()
}
