/*
Implement a worker pool (3 workers) where:

Each worker processes jobs from a channel.

You use a context.WithTimeout of 2 seconds.

If jobs take longer, workers should stop gracefully.
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func workerFunc(wg *sync.WaitGroup, ctx context.Context, ch chan int) {
	defer wg.Done()
	fmt.Println("Worker created")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context got cancelled, returning from worker")
			return
		case v, ok := <-ch:
			if !ok {
				return
			}
			fmt.Println("Got a job to process on", v)
			time.Sleep(1 * time.Second)
		}
	}

}
func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		count := 20
		for i := 0; i < count; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Context got cancelled, no more job needs to get executed")
				return
			case ch <- i:
				fmt.Println("Main: sent job", i)
				time.Sleep(1 * time.Second)
			}
		}
		close(ch)
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go workerFunc(&wg, ctx, ch)
	}

	wg.Wait()
	fmt.Println("Main completed...")
}
