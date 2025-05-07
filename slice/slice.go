package main

import "fmt"

func main() {
	// list that contains multiple items of same type
	// no need to declare the size of list in advance unlike array
	var mySlice []string
	fmt.Println("My slice", mySlice, len(mySlice), cap(mySlice))
	// let's try to assign some value to it.
	// list is empty i.e. reference of an array so need to first intialize it otherwise it will panic
	mySlice = make([]string, 1)
	mySlice[0] = "adarsh"
	fmt.Println("My slice", mySlice, len(mySlice), cap(mySlice))
	// below will again panic as the size is 1
	/*mySlice[1] = "adarsh"
	fmt.Println("My slice", mySlice, len(mySlice), cap(mySlice))
	*/
	// we can append the new element even if initial size was 1
	mySlice = append(mySlice, "yadav")
	fmt.Println("My slice", mySlice, len(mySlice), cap(mySlice))
}
