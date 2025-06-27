package main

import (
	"fmt"
	"sync"
)

// ✅ Sends squares of 1–5 to its channel
func square(myChan chan int) {
	for i := 1; i < 6; i++ {
		myChan <- i * i // send square of i
	}
	close(myChan) // very important: close channel after all values sent
}

// ✅ Sends doubles of 6–10 to its channel
func double(myChan chan int) {
	for i := 6; i < 11; i++ {
		myChan <- i * 2 // send double of i
	}
	close(myChan) // close the channel after done
}

// ✅ Fan-in function: Merges multiple input channels into one output channel
func mergedChanRoutine(ch ...chan int) chan int {
	mergedChan := make(chan int) // This will carry merged values from all input channels
	wg := sync.WaitGroup{}       // WaitGroup to wait for all goroutines to finish

	// output() copies values from input channel to mergedChan
	output := func(c chan int) {
		for v := range c { // Read from input channel
			mergedChan <- v // Write into merged channel
		}
		wg.Done() // Done when this input channel is fully read and closed
	}

	wg.Add(len(ch)) // Add number of channels to the waitgroup

	for _, v := range ch {
		go output(v) // Start one goroutine per input channel to fan-in
	}

	// ✅ Once all output() goroutines are done, close the merged channel
	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	return mergedChan // Return the merged output channel
}

func main() {
	chan1 := make(chan int) // Will get squares of 1–5
	chan2 := make(chan int) // Will get doubles of 6–10

	// Start workers that send data to chan1 and chan2
	go square(chan1)
	go double(chan2)

	// Fan-in both channels into a single merged channel
	mergedChan := mergedChanRoutine(chan1, chan2)

	// ✅ Read and print all values from merged channel
	for v := range mergedChan {
		fmt.Println(v)
	}

	fmt.Println("Main completed")
}
