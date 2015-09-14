package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6))
	fmt.Println(add(1, 3, 4))
	fmt.Println(add())
	// Passing a slice of ints.
	xs := []int{100, 200, 300}
	fmt.Println(add(xs...))
}
