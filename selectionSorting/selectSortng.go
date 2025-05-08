package main

import "fmt"

/*
| Case       | Time  |
| ---------- | ----- |
| Best Case  | O(n²) |
| Worst Case | O(n²) |
| Space      | O(1)  |

🧠 When to Use It?
When memory is very limited (uses no extra space)

Educational purposes

Not used in real apps for large data (slow)
*/

func main() {
	unsortedSlice := []int{29, 14, 10, 37, 13}
	fmt.Println(unsortedSlice)
	// selects the smallest element from the unsorted part and swaps it with the first unsorted element
	for i := 0; i < len(unsortedSlice)-1; i++ {
		// select the smallest element from the unsorted array.
		// assume the minNoIndex = i
		minIdexNo := i
		for j := i + 1; j < len(unsortedSlice); j++ {
			if unsortedSlice[j] < unsortedSlice[minIdexNo] {
				// if any other smaller no. is found then store the index of that element
				minIdexNo = j
			}
		}
		// swap the assumed minNo i.e. ith element with the real smallest no.
		unsortedSlice[i], unsortedSlice[minIdexNo] = unsortedSlice[minIdexNo], unsortedSlice[i]
	}
	fmt.Println(unsortedSlice)
}
