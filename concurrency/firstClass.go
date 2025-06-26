/*
Assignment:
✅ Write a program with a goroutine that prints squares of numbers from 1 to 5.
✅ Add time.Sleep to make sure goroutine completes.
❌ Remove time.Sleep → Observe what happens (goroutine may not run).
✅ Use multiple go statements to run 2+ goroutines simultaneously.
🧪 Try printing runtime.NumGoroutine() to see how many are running.
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

func squares() {
	for i := range 10 {
		fmt.Println(i * i)
	}
}
func main() {
	go squares()
	go squares()
	go squares()
	fmt.Println("No of goroutine", runtime.NumGoroutine(), "No. of CPU", runtime.NumCPU())
	time.Sleep(1 * time.Second)
	fmt.Println("main done")
}
