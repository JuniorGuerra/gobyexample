package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("unitint", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp", s, len(s), "cap:", cap(s))
	// cap is max size, len is the actual size

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("adp: ", s)
	fmt.Println("len: ", len(s), "cap: ", cap(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copy", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	l = s[:5]
	fmt.Println("sl2", l)

	l = s[2:]
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	usesSlices()
}

func usesSlices() {
	slice := []int{10, 2, 12, 3, 4, 512, 123, 12}
	slices.Sort(slice)

	fmt.Println(slice)

	if slices.IsSorted(slice) {
		fmt.Println("Is sorted")
	}
	var slice2 = make([]int, len(slice))
	copy(slice2, slice)
	slices.SortFunc(slice2, func(a, b int) int {
		if a < b {
			return b
		}
		return a
	})
	if slices.Equal(slice, slice2) {
		fmt.Println("there are sorted")
	}
	slices.Delete(slice2, 1, 5)

	slices.Reverse(slice2)
	
	fmt.Println(slice2)
}
