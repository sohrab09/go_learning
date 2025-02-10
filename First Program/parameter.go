package main

import "fmt"

func add(a int, b int) { // a and b are parameters
	c := a + b
	fmt.Println(c)
}

func main() {
	add(10, 20) // 10 and 20 are arguments
}
