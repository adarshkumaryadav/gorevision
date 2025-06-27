/*📌 Goal:
Generate numbers 1–10

Square them

Double the result

Print in main

----
Write three functions:

1️⃣ generate(chan int)
Sends 1 to 10 to chan int, then closes the channel

2️⃣ square(chan int) -> chan int
Receives from input channel

Sends n * n to a new output channel

Closes output channel after done

3️⃣ double(chan int) -> chan int
Same idea: receives → doubles → sends → closes

4️⃣ main
Ranges over final output channel and prints values
*/

package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	go func() {
		for i := 1; i < 11; i++ {
			chan1 <- i
		}
		close(chan1)
	}()
	go func() {
		for v := range chan1 {
			chan2 <- v * v
		}
		close(chan2)
	}()
	go func() {
		for v := range chan2 {
			chan3 <- v * 2
		}
		close(chan3)
	}()
	sum := 0
	for v := range chan3 {
		sum += v
		fmt.Println(v)
	}

}
