package main

import (
	"fmt"
	"time"
)

var chan1 = make(chan int)
var chan2 = make(chan int)

func firstGoRoutine() {
	time.Sleep(1 * time.Second)
	chan1 <- 1
}
func secondGoRoutine() {
	time.Sleep(2 * time.Second)
	chan2 <- 2
}
func main() {
	go firstGoRoutine()
	go secondGoRoutine()
	select {
	case v := <-chan1:
		fmt.Println("Read on first chan: ", v)
	case v := <-chan2:
		fmt.Println("Read on second chan: ", v)
	// default:
	//fmt.Println("Default") // as both goroutine has sent the data default will get executed
	case <-time.After(3 * time.Second):
		fmt.Println("Waited for 3 seconds")
	}

	fmt.Println("Main routine completed")

}
