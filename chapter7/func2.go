package main

import "fmt"

func multiple_return() (int, int) {
	return 5, 6
}

func main() {
	x, y := multiple_return()
	fmt.Println(x, y)
}
