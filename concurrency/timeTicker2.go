/*
🎯 Assignment: Rate-Limited Worker Pool
✅ Objective:
You have 10 jobs to process.

You have 3 workers.

Only 1 job should be processed every second (rate limit).

Each job just prints: Worker X processed job Y at HH:MM:SS
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobChain chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for v := range jobChain {
		fmt.Printf("Worker id: %d ,Got a job to process on %d , %v\n", id, v, time.Now().Format("15:04:05"))
	}
}
func main() {
	jobChain := make(chan int)
	wg := sync.WaitGroup{}
	go func(chan int) {
		ticker := time.Tick(1 * time.Second)
		for i := 1; i < 11; i++ {
			<-ticker
			jobChain <- i
		}
		close(jobChain)
	}(jobChain)

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go worker(jobChain, &wg, i)
	}
	wg.Wait()
	fmt.Println("Main completed")
}
