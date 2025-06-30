/*
🔜 Let’s Begin With: Rate Limiting
🔹 Scenario:
We have an API client that can only make 1 request per second.
 We want to simulate and control this using a Go program.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(1 * time.Second)
	for i := 0; i < 5; i++ {
		// wait till ticker got triggered that will eventually be triggered after each 1 seconds
		<-ticker
		fmt.Println("Going to send the request ...")
	}
}
