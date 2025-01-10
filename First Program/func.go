package main

import "fmt"

func add(a int, b int) int { // return a single value
	return a + b
}

func getNumbers(a int, b int) (int, int) { // return multiple values
	sum := a + b

	mul := a * b

	return sum, mul

}

func main() {
	a := 10
	b := 20
	c := add(a, b)
	fmt.Println(c)

	d, e := getNumbers(a, b)
	fmt.Println(d, e)
}
