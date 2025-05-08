package main

import "fmt"

/*
⏱ Time & Space Complexity
Case	Time
Best Case	O(n) ✅ (with swapped optimization)
Worst Case	O(n²) ❌
Space	O(1) ✅

🧠 When to Use It?
For educational purpose

# Very small data sets

When you want to teach comparison/swapping logic
*/
func main() {
	// unsortedSlice := []int{10, 999, 5, 34, 123, 1, 0, 23}
	// already Sorted
	unsortedSlice := []int{10, 11, 16, 20, 123, 1222, 22222, 2322222}
	len := len(unsortedSlice)
	for i := 0; i < len-1; i++ {
		sortingNeeded := false
		fmt.Println("Inside ith loop", unsortedSlice, i)

		for j := 0; j < len-1; j++ {
			// we will try to move the first and then another biggest no. at the end once this inner loops completed
			fmt.Println("Comparing", unsortedSlice[j], unsortedSlice[j+1])
			if unsortedSlice[j] > unsortedSlice[j+1] {
				// at any point in any iteration if everything was sorted then none of the jth will be bigger than j+1
				// it means it's sorted, no need to iterate any more
				sortingNeeded = true
				unsortedSlice[j], unsortedSlice[j+1] = unsortedSlice[j+1], unsortedSlice[j]
			}
		}
		if !sortingNeeded {
			break
		}
		fmt.Println(unsortedSlice)
	}
	fmt.Println(unsortedSlice)
}
