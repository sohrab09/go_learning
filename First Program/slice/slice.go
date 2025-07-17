package main

import "fmt"

func changeSlice(a []int) []int {
	a[0] = 10
	a = append(a, 11)
	return a
}

func main() {
	// arr := [6]string{"This", "is", "a", "slice", "example", "test"}
	// fmt.Println(arr)

	// sl := arr[1:3]
	// fmt.Println(sl)

	// s1 := sl[1:2]
	// fmt.Println(s1)
	// fmt.Println(len(s1)) // know the length of slice
	// fmt.Println(cap(s1)) // know the capacity of slice
	// fmt.Println(s1[0])   // know the first element of slice

	// // slice literal
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)

	// make function
	// slice1 := make([]int, 5)
	// fmt.Println(slice1)

	// empty slice or nil slice
	// var s []int

	// s = append(s, 1, 2, 3, 4, 5)

	// fmt.Println(s) // nil slice

	// var x []int
	// x = append(x, 1)
	// x = append(x, 2)
	// x = append(x, 3)

	// y := x

	// x = append(x, 4)

	// y = append(y, 5)

	// x[0] = 10

	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)
	fmt.Println("x:", x)
	fmt.Println("y:", y)

}

/*
	2 types Phases:
		1. Compilation Phase (compile time)
		2. Execution Phase (run time)
*/

/*
	স্লাইস এর মধ্যে মেইন ৩ টা জিনিস থাকে:
		1. Pointer: স্লাইস এর প্রথম এলিমেন্টের ঠিকানা
		2. Length: স্লাইস এর মধ্যে কয়টা এলিমেন্ট আছে
		3. Capacity: স্লাইস এর মধ্যে কয়টা এলিমেন্ট আছে
*/

/*
	We can create slice with 3 ways:
		1. Slice from existing array
		2. Slice from a slice
		3. Using slice literal
		4. Using make function with len
		5. Using make function with len and cap
		6. empty slice or nil slice
		7. slice underlying array rule => 1024 -> 100% increase, 1024 < -> 25% increase
*/
