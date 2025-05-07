package main

import "fmt"

// binary searching means that we are dividing our search operation in half
// it is applied on the sorted DS like sorted array or slice
// it's complexity is o(log n) as we are deducting the list/2 in each iteration

func main() {
	sortedArray := [10]int{23, 56, 89, 100, 120, 160, 176, 190, 234, 333}
	// first we have to find the mid element
	// if mid is not my target, we will try to shift the low and high or left or right
	// for now at the starting consider that the low is the first and hight is the last element
	low, high := 0, len(sortedArray)-1
	target := 190
	for low <= high {
		// find the middle point of both cornergi
		mid := (high + low) / 2
		if target == sortedArray[mid] {
			fmt.Println("Target is at ", mid)
			break
		}
		if target > sortedArray[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

}
