package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// myChan := make(chan int)
	/* this is go into deadlock
	// myChan <- 76 // can'r send the data if there is no reciever
	// result := <-myChan // can't do this as there should be a reciever in parallel in another routine at the same time
	fmt.Println("Main routine is done")
	*/

	/* even this  one will go into deadlock as the reciever is not available before sender start sending the data
	myChan <- 76
	// anonymous go routine
	go func() {
		fmt.Println(<-myChan)
	}()
	time.Sleep(2 * time.Second)
	*/

	// but below one will work
	// anonymous go routine
	/* go func() {
		fmt.Println(<-myChan)
	}()
	myChan <- 76

	time.Sleep(2 * time.Second)
	*/

	// now let's try our buffered chan
	bufferChan := make(chan int, 1)
	/* this will work as the chan has a buffer of cap 1 so sender can send the data over chan even if there is no reciever at the moment
	bufferChan <- 76
	fmt.Println(<-bufferChan)
	*/

	/*
		bufferChan <- 76
		// below will not work as cap of buffer is one only
		bufferChan <- 77
	*/
	bufferChan <- 76
	// will be in deadlock as there is no sender
	// if at line no 44 one data is sent it will read but howevr after that again deadlock
	/*for value := range bufferChan {
		fmt.Println(value)
	}
	*/

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg2 := sync.WaitGroup{}
		for i := range 100 {
			wg2.Add(1)
			go func() {
				defer wg2.Done()
				time.Sleep(1 * time.Second)
				bufferChan <- i
			}()
		}
		wg2.Wait()
		// chan is still open
		close(bufferChan)
		wg.Done()
	}()

	// reciever: rec can't close a chan
	wg.Add(1)
	go func() {
		// even this will go in deadlock as reciever is still looking for val
		for i := range bufferChan {
			fmt.Println(i)
		}
		wg.Done()
	}()
	wg.Wait()
}
