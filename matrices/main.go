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

func prodMA() {
	var a [3][3]int = [3][3]int{
		{3, 1, -2}, {0, 4, 2}, {7, 5, 1},
	}
	var b [3][3]int = [3][3]int{
		{-1, 0, 8}, {3, 6, 9}, {0, 0, 3},
	}
	var c [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = 0
			for k := 0; k < 3; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	fmt.Println(c)
}

func multidimensionalArray() {
	// staticMa()
	prodMA()
}

func main() {
	multidimensionalArray()
}
