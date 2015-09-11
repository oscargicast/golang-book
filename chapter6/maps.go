package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key1"] = 10
	x["key2"] = 20
	fmt.Println(x)
	fmt.Println(x["key2"])
}
