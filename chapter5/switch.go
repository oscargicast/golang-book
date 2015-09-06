package main

import "fmt"

func main() {
	var i int
	for {
		fmt.Print("Enter number from 0 to 6: ")
		fmt.Scanf("%d", &i)
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		default:
			fmt.Println("Bad boy, try again ;)")
		}
	}
}
