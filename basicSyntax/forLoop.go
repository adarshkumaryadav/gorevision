package main

import "fmt"

func main() {
	// iterating array
	var myArray = [5]string{"1", "2", "3", "4", "5"}
	for index, value := range myArray {
		fmt.Printf("At index %v , value is %v\n", index, value)
	}
	// infinite loop
	var i = 0
	for {
		fmt.Println("Iteration", i)
		i++
		if i == 10 {
			// breaking the loops
			break
		}
	}
}
