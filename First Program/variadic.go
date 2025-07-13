package main

import "fmt"

// variadic function

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}

func main() {
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
