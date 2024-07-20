package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

	val, ok := m["k2"]
	fmt.Println("val: ", val, " prs: ", ok)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	mapsUse()
}

func mapsUse() {
	m := map[string]string{
		"monday":   "monday",
		"sunday":   "sunday",
		"saturday": "weekday",
	}

	z := maps.Clone(m)
	fmt.Println(z)

	maps.DeleteFunc(m, func(v, v1 string) bool {
		if v == v1 {
			return true
		}
		return false
	})

	fmt.Println(m)
}
