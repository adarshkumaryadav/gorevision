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

	tempArray := [10]int{1, 2, 3, 56, 23, 19, 0, 11, 12, 19}
	var tempArray2 [10]int
	for i, v := range tempArray {
		if i == 0 {
			tempArray2[0] = tempArray[0]
			continue
		}
		tempArray2[i] = v + tempArray2[i-1]
	}

	fmt.Println(tempArray2)
	// q time o(q)
	for {
		var left, right int
		fmt.Println("Enter left")
		fmt.Scan(&left)
		fmt.Println("Right")
		fmt.Scan(&right)
		if right < left {
			fmt.Println("Error...")
			return
		}
		//var sum int
		//o(r-l)
		// 10-0 = 10 = n
		/*for i := left; i < right; i++ {
			sum += tempArray[i]
		}*/
		sum := (tempArray2[right-1]) - (tempArray2[left-2])
		fmt.Println("sum", sum)
	}
	// time = o(q*n) // space o(1)
	// requirement = time = o(q) ... spacce = o(n)
}
