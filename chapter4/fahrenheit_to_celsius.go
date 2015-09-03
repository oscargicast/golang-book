package main

import "fmt"

func main() {
	var f float64
	fmt.Print("Enter Temperature in C: ")
	fmt.Scanf("%f", &f)
	c := (f - 32) * 5 / 9
	fmt.Print("Temperature in F: ")
	fmt.Println(c)
}
