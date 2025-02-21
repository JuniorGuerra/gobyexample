package main

import "fmt"

func sum(nums ...int) {
	fmt.Println(nums, "")

	fmt.Printf("the type is: %T\n", nums)

	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2, 3, 4)
	sum(1, 2, 3, 4, 5)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
