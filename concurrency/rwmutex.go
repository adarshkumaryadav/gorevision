/*📌 When to use RWMutex?
Use it when:

You have frequent reads and rare writes

You want to allow multiple reads in parallel, but writes should be exclusive
| Lock Type | Purpose                        | Allows             | Blocks              |
| --------- | ------------------------------ | ------------------ | ------------------- |
| `RLock()` | Read-only access (multiple OK) | ✅ Multiple readers | ❌ Writers           |
| `Lock()`  | Write access (exclusive)       | ❌ Only one writer  | ❌ Readers & Writers |
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	myMap := make(map[string]int)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) { // only writing in these routines
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			fmt.Println("update the value")
			myMap["x"]++
		}(10000 + i)
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) { // only reading in these routines
			defer wg.Done()
			mu.RLock()
			defer mu.RUnlock()
			fmt.Println(myMap["x"])
		}(i)
	}

	wg.Wait()
	fmt.Println("Main completed")
	fmt.Println(myMap["x"])
}
