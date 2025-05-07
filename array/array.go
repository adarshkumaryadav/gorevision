package main

import "fmt"

func main() {
	var myArray = [10]string{}
	// both len and cap will be fixed and it will be in this case
	fmt.Println("myArray", myArray, len(myArray), cap(myArray))
	myArray[1] = "adarsh"
	// not a valid operation
	// myArray = []string{"", ""}
	var anotherArray = [5]string{}
	// not allowed as size is diff
	// myArray = anotherArray
	fmt.Println(anotherArray)
	anotherArray2 := [10]string{}
	// this is allowed
	myArray = anotherArray2
	fmt.Println(myArray, anotherArray2)
}
