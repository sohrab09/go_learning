package main

import "fmt"

func main() {

	// if else condition

	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	// switch case

	y := 30
	switch y {
	case 30:
		fmt.Println("y is 30")
	case 20:
		fmt.Println("y is 20")
	default:
		fmt.Println("y is neither 30 nor 20")
	}
}
