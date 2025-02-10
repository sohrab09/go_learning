package main

import "fmt"

var a = 10
var b = 20

func add(x int, y int) int {
	z := x + y
	fmt.Println(z)
	return z
}

func main() {

	p := 30
	q := 40

	add(p, q)

	add(a, b)

	add(a, p)

	// add(b, z)
}
