package main

import "fmt"

/*
🧠 Intuition:
Start from the second element.

Compare it with the left side (sorted side).

Shift bigger elements to the right to make space.

Insert the current element at the correct position.

🔢 Dry Run (Example)
Unsorted Array: [5, 2, 4, 6, 1, 3]

Step-by-step:

2 compared with 5 → insert before → [2, 5, 4, 6, 1, 3]

4 → move 5 → insert at correct → [2, 4, 5, 6, 1, 3]

6 → already in place

1 → move 6, 5, 4, 2 → insert at 0 → [1, 2, 4, 5, 6, 3]

3 → move 6, 5, 4 → insert after 2 → [1, 2, 3, 4, 5, 6]

⏱️ Time & Space Complexity
Case	Time Complexity
Best Case	O(n) – already sorted
Average	O(n²)
Worst Case	O(n²) – reverse sorted

Space: O(1) – in-place
*/
func main() {
	// unsorted array
	unsortedSlice := []int{23, 43, 2, 56, 4, 3, 88, 2}
	for i := 1; i < len(unsortedSlice); i++ {
		// pick one element at a time and place it into it's right position after comparing it with prev index
		pickedNo := unsortedSlice[i]
		j := i - 1
		for ; j >= 0 && unsortedSlice[j] > pickedNo; j-- {

			unsortedSlice[j+1] = unsortedSlice[j]

		}
		unsortedSlice[j+1] = pickedNo
	}
	fmt.Println(unsortedSlice)
}
