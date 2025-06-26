package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			chan1 <- "From A"
		}
	}()

	go func() {
		for {
			time.Sleep(800 * time.Millisecond)
			chan2 <- "From B"
		}
	}()
	timeOut := time.After(5 * time.Second)
Loop:
	for {
		select {
		case v := <-chan1:
			fmt.Println(v)
		case v := <-chan2:
			fmt.Println(v)
		// case <-time.After(5 * time.Second): // this is being starte in each select so it will never be triggered as gorutine are writing data before 5 seconds
		case <-timeOut:
			break Loop
		}
	}

	fmt.Println("Main completed")
}
