/*
✅ Assignment:
Create a job channel (chan int)

Create a result channel (chan int)

Launch 3 goroutines as workers:

# Each should take a number from the job channel

# Square the number

# Send it to the result channel

# In main, send 10 numbers to the job channel

# Close the job channel once sending is done

# Use a WaitGroup to wait for workers to finish

Range over the result channel and print the results
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	job := make(chan int)
	result := make(chan int)
	sum := 0
	wg := sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range job {
				result <- v * v
			}
		}()
	}
	// need to close the result chan as well once workers are done with updating the results
	go func() {
		wg.Wait()
		close(result)
	}() // this will be done once all above workers are done
	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for v := range result { // this will wait until above anonymous routine is not closing the result chan
			sum += v
			fmt.Println(sum)
		}
	}()
	for i := 1; i < 11; i++ {
		job <- i
	}
	close(job)

	wg2.Wait()

	fmt.Println("total sum is: ", sum)
}
