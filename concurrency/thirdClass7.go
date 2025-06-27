/*
📋 Assignment Details:
Create:

jobChan → sends numbers 1 to 30

resultChan → receives results like: Number: 17 is Prime or Number: 18 is NOT Prime

# Fixed number of 3 workers

Each worker:

# Picks a number from jobChan

Checks if it’s prime (write a small isPrime(n int) function)

# Sends the message string to resultChan

Use:

sync.WaitGroup to wait until all workers finish

# Close resultChan after all work is done

Print results in main by ranging over resultChan
*/
package main

import (
	"fmt"
	"sync"
)

func isPrime(wg *sync.WaitGroup, jobChain chan int, resultChain chan string, id int) {
	defer wg.Done()
	for v := range jobChain {
		if primeCheck(v) {
			resultChain <- fmt.Sprintf("%d is a prime number, job id: %d", v, id)
		} else {
			resultChain <- fmt.Sprintf("%d is not a prime number, job id: %d", v, id)
		}
	}
}

func primeCheck(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	wg := sync.WaitGroup{}
	jobChain := make(chan int)
	resultChain := make(chan string)
	// create the job
	go func() {
		for i := 0; i <= 100; i++ {
			jobChain <- i
		}
		close(jobChain)
	}()

	// worker got created, this will execute in concurrent
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go isPrime(&wg, jobChain, resultChain, i)
	}
	// this should get completed other wise deadlock will occur as main is ranging over resultChain
	go func() {
		wg.Wait()
		close(resultChain)
	}()

	// now it's time to get the result and print them.
	for v := range resultChain {
		fmt.Println(v)
	}

}
