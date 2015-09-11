package main

import "fmt"

func main() {
	slice0 := make([]float64, 6, 10)
	slice0[0] = 6
	slice0[1] = 1
	fmt.Println(slice0)
	// Append.
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	// Copy.
	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)
}
