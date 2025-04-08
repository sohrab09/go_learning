package main

import "fmt"

// pass by value example
// In this example, we pass the array itself to the function. This means that a copy of the array is made, and any changes made to the array inside the function will not affect the original array.

/*
	func print(numbers [5]int) {
		fmt.Println(numbers)
	}

	func main() {
		arr := [5]int{1, 2, 3, 4, 5}
		print(arr)
	}
*/

// pass by reference example
// In this example, we pass a pointer to the array instead of the array itself.
/*
	func print(numbers *[5]int) {
		fmt.Println(numbers)
	}

	func main() {

		arr := [5]int{1, 2, 3, 4, 5}
		print(&arr)
	}
*/

// example with structs
type Person struct {
	Name string
	Age  int
}

func main() {
	info := Person{
		Name: "Nahid",
		Age:  28,
	}

	i := &info      // i is a pointer to info
	fmt.Println(*i) // prints the value of info
}
