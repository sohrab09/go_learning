package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func printUserDetails(usr User) {
	fmt.Println("Name: ", usr.Name)
	fmt.Println("Age: ", usr.Age)
}

func main() {
	var user1 User
	user1 = User{ // it's another name is instantiate
		Name: "John",
		Age:  30,
	}
	printUserDetails(user1)

	user2 := User{ // instance or object
		Name: "Doe",
		Age:  40,
	}
	printUserDetails(user2)
}
