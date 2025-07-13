package main

import "fmt"

// defer statement example
func a() {
	i := 0
	fmt.Println("first:", i)

	defer fmt.Println("second:", i)

	i = i + 1

	fmt.Println("third:", i)

	defer fmt.Println("fourth:", i)

	return
}

// Named return values
func sum(a int, b int) (result int) {
	result = a + b
	return
}

// defer functions with named return values

func calculate() (result int) {
	fmt.Println("Calculating...", result)
	show := func() {
		result = result + 10
		fmt.Println("defer:", result)
	}
	defer show()
	result = 5
	return
}

func main() {
	a()
	res := sum(5, 10)
	fmt.Println("Result:", res)
	calculate()
}


/*
	** rules of named return values:
		1. all codes execute
		2. defer function store kora hobe magic box e
		3. return -> all defer function execute korbe
		4. return korbe named variable gular values

	** rules of just return values:
		1. all codes execute
		2. defer function store kora hobe magic box e
		3. return values are evaluated at this time (store the return value)
		4. all defer function execute korbe
/*