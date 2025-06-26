package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		chan1 <- 1
	}()
	select {
	case <-chan1:
		fmt.Println("Rcvd value on chan1")
	case <-time.After(2 * time.Second):
		fmt.Println("No value rcvd ...waited for 2 seconds")
	}

	chan2 := make(chan int, 2)
Loop:
	for {
		select {
		case chan2 <- time.Now().Local().Second():
			fmt.Println("Vlaue sent value send", time.Now().Local().Second())
		default:
			fmt.Println("Channel full")
			break Loop
		}
	}
	fmt.Println("Main completed")
}
