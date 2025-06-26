/*
	1️⃣ Try sending to a closed channel

Expected: panic: send on closed channel

2️⃣ Try closing the same channel twice
Expected: panic: close of closed channel

3️⃣ Try receiving from a closed channel
Expected: You won’t panic — you’ll get:

# Zero value

ok = false
*/
package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1
	close(c)
	// c <- 1 // panic as we can't write to a close channel
	// close(c) // panic..can't close a chan twice
	v, ok := <-c       // zeor value of chan type and ok will be false
	fmt.Println(v, ok) // if buffered chan even if chan is closed 1 value sent earlier will be there
	// note this time ok will be true so it means that ok doesn't gurantee that chan is close or not in buffer chan case
}
