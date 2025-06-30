package main

import (
	"sync"
	"testing"
)

func BenchmarkWithMutex(b *testing.B) {
	var mu sync.Mutex
	counter := 0
	for i := 0; i < b.N; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func BenchmarkWithChannel(b *testing.B) {
	ch := make(chan int, 1)
	ch <- 0
	for i := 0; i < b.N; i++ {
		val := <-ch
		val++
		ch <- val
	}
}
