package main

import "fmt"

func makeEventGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) { // Naming the returning type.
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEventGenerator()
	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println(nextEven()) //4
}
