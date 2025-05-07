package main

import "fmt"

func main() {
	type myStruct struct {
		name string
		age  int
	}
	var user myStruct
	fmt.Println("Enter your name")
	fmt.Scan(&user.name)
	fmt.Println("enter age")
	fmt.Scan(&user.age)
	fmt.Println("Here is your data", user)
}
