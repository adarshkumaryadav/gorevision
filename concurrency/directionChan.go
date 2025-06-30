package main

import (
	"fmt"
)

func sender(myChan chan<- int) {
	for i := 1; i < 6; i++ {
		myChan <- i
	}
	close(myChan)
}
func reciever(myChan <-chan int) {
	for v := range myChan {
		fmt.Println(v)
	}
}
func main() {
	myChan := make(chan int)
	go sender(myChan)
	reciever(myChan)

	//time.Sleep(1 * time.Second)
}
