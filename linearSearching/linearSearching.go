package main

import "fmt"

// bubble searching is used for unsorted case
// if i have an array/slice that is unsorted
// this is easiest one and having o(n) time complexity
func main() {
	myUnsortedSlice := []int{1, 2, 56, 23, 234, 5, 78, 98, 32, 1}
	// now i have to search 78 out of these data stored in slice
	// iterate over each and every element with for loop
	// in worst case it can be found at the last element so complexity will be o(n) as n is the no. of iteration
	for i, v := range myUnsortedSlice {
		if v == 78 {
			fmt.Printf("Found element at %v index \n", i)
			break
		}
	}
}
