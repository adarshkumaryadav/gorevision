/*
Unbuffered Ping-Pong

Create 2 goroutines: ping and pong.

Use two unbuffered channels so they send and receive “ping”/“pong” back and forth N times.

Buffered Generator

Make a buffered channel of capacity 5.

Launch a goroutine that generates numbers 1–10 and sends them into the channel, then closes it.

In main, range over the channel to print each number.

Deadlock Exploration

Try sending 6 items into a buffered channel with capacity 5 (without closes/receives).

Observe the panic/deadlock.

Channel Close Edge-Case

Close a channel, then try sending → observe runtime panic.

Try closing the channel twice → observe panic.

Optional:

Print runtime.NumGoroutine() before and after channel operations to see how many are active.
*/

package main

import (
	"fmt"
	"sync"
)

func ping(wg *sync.WaitGroup, x, y chan string) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		x <- "ping"
	}
	close(x)
	for value := range y {
		fmt.Println(value)
	}
}

func pong(wg *sync.WaitGroup, x, y chan string) {
	defer wg.Done()
	for value := range x {
		fmt.Println(value)
	}
	for i := 0; i < 10; i++ {
		y <- "pong"
	}
	close(y)
}

func ping2(wg *sync.WaitGroup, x chan string) {
	defer wg.Done()
	count := 0
	x <- "ping2"
	for v := range x { // issue is that same go routine write/send and can read
		fmt.Println(v, "reading From ping")
		x <- "ping2"
		count++
		if count == 10 {
			close(x)
		}
	}
}
func pong2(wg *sync.WaitGroup, x chan string) {
	defer wg.Done()
	for v := range x {
		fmt.Println(v, "reading from pong")
		x <- "pong2"
	}
}
func main() {
	wg := sync.WaitGroup{}

	/*
		firstChan := make(chan string)
		secondChan := make(chan string)

		wg.Add(2)
		// ping ping .....pong pong ....
		go ping(&wg, firstChan, secondChan)
		go pong(&wg, firstChan, secondChan)

		// ping pong ping pong ....
		thirdChan := make(chan string)
		wg.Add(2)
		go ping2(&wg, thirdChan)
		go pong2(&wg, thirdChan)
		wg.Wait()
	*/
	fourthChan := make(chan string)
	wg.Add(2)
	go ping3(&wg, fourthChan)
	go pong3(&wg, fourthChan)
	fourthChan <- "ping4"
	wg.Wait()
}

func ping3(w *sync.WaitGroup, x chan string) {
	defer w.Done()
	for {
		v, ok := <-x
		if !ok {
			return
		}
		if v == "ping4" {
			fmt.Println("ping4")
			x <- "pong4"
		}
	}

}

func pong3(w *sync.WaitGroup, x chan string) {
	defer w.Done()
	count := 0
	for {
		v := <-x
		if v == "pong4" {
			fmt.Println(v)
			count++
		}
		if count > 6 {
			close(x)
			return
		}
		x <- "ping4"
	}
}
