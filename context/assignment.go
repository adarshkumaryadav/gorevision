package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func consumerFunc(ctx context.Context, job chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context got cancelled or timeout")
			return
		case v, ok := <-job:
			if !ok {
				fmt.Println("Chan got closed, returning")
				return
			}
			time.Sleep(1 * time.Second)
			fmt.Println("Found a job to process some op on ", v)
		}

	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	jobChain := make(chan int)

	// creating the consumer
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go consumerFunc(ctx, jobChain, &wg)
	}

	// job sending
	go func(ctx context.Context) {
		// no need to add wg to wait for this to get compelted as the consumer is
		// already being waited to get compelted so until consumer is readign from chain this will be live
		defer close(jobChain)
		for i := 1; i < 101; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Context got cancelled or timeout")
				return
			case jobChain <- i:
				fmt.Println("Send a job for ", i)
			}
		}
	}(ctx)
	time.AfterFunc(2*time.Second, cancel)
	fmt.Println("Going to wait for all routine to get completed")
	wg.Wait()
	fmt.Println("Main competed")
}
