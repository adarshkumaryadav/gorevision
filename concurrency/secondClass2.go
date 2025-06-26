/*
Create a buffered channel with capacity 5.
Launch a goroutine that sends numbers 1–10 into the channel.
After sending, close the channel.
In main(), use range to read and print all values from the channel.
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	bufferChan := make(chan int, 5)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			bufferChan <- i
		}
		close(bufferChan)
	}()
	fmt.Println("Active goroutines:", runtime.NumGoroutine())

	for v := range bufferChan {
		fmt.Println(v)
	}
	//bufferChan <- 9
	close(bufferChan)
	// The wg.Wait() is not strictly needed, because once the channel is closed, range will automatically exit. By that time, the goroutine is anyway done.
	wg.Wait()
}
