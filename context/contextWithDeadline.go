package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	timeVar := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), timeVar)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context got cancelled, retuning")
				return
			default:
				fmt.Println("Doing some job")
				time.Sleep(1 * time.Second)
			}

		}
	}(ctx)
	wg.Wait()
	fmt.Println("Main completed")
}
