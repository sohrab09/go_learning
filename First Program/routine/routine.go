package main

import (
	"fmt"
	"time"
)

var n = 10

const m = 11

func printHello(num int) {
	fmt.Printf("Hello, World! %d\n", num)
}

func main() {
	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

	fmt.Println("Main function is running", n, m)
	time.Sleep(5 * time.Second)
}
