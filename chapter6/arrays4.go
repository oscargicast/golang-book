package main

import "fmt"

func main() {
	//var x = [5]float64{98, 93, 7, 82, 83}
	x := [5]float64{98, 93, 7, 82, 83}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
