// Internal Memory

package main

import "fmt"

const a = 10

var p = 100

func call() {
	test := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	test(5, 6)

	test(p, a)
}

func main() {
	call()

	fmt.Println(a)
}

func init() {
	fmt.Println("Hello")
}
