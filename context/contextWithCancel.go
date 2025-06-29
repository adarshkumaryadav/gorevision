package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
Loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context got cancelled")
			// if not breaking this will keep hitting this case until main exist
			break Loop

		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Default")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
	fmt.Println("Main completed")
}
