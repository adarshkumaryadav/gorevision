/*
🧪 Small Task For You (Do this before we move on)
🔁 Try implementing a map[string]int counter:

10 goroutines increment a counter

# Use sync.Mutex to protect the map

Print final map safely
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	myMap := make(map[string]int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("Id: ", id)
			mu.Lock()

			defer mu.Unlock()
			myMap["x"]++
		}(i)
	}
	wg.Wait()
	fmt.Println(myMap["x"])
}
