

Here's a **step-by-step roadmap** of topics to cover, designed for real-world development and interview readiness.

---

## 🧭 Go Concurrency Mastery Roadmap

### 🟢 1. **Goroutines**

* What are goroutines?
* How they differ from OS threads
* How to start a goroutine
* Common pitfalls: forgetting to `Wait`, goroutine leaks

### 🟢 2. **Channels (Unbuffered & Buffered)**

* `chan` syntax
* Blocking behavior of send/receive
* Buffered vs unbuffered channels
* Deadlocks and how to avoid them
* Ranging over channels
* Closing channels properly

### 🟡 3. **`select` Statement**

* Multiplexing channel operations
* `select` with default (non-blocking ops)
* `select` with timeouts (`time.After`)
* `select` with cancellation

### 🟡 4. **Wait Groups (`sync.WaitGroup`)**

* Coordinating multiple goroutines
* Waiting for completion
* Common mistake: calling `Done()` too early or too late

### 🟡 5. **Mutexes and RWMutexes (`sync.Mutex`, `sync.RWMutex`)**

* Mutual exclusion for shared data
* Read/Write locking
* Defer pattern (`defer mu.Unlock()`)

### 🟡 6. **Once and Atomic Operations**

* `sync.Once` for one-time initialization
* `sync/atomic` for lock-free counters and flags

---

## 🔵 Intermediate Concepts

### 🔵 7. **Context Package (`context.Context`)**

* Context with cancellation
* Deadlines and timeouts
* Passing context across API boundaries
* `context.WithCancel`, `WithTimeout`, `WithDeadline`, `WithValue`

### 🔵 8. **Worker Pool Pattern**

* Creating and managing a pool of goroutines
* Channel-based job queues
* Graceful shutdown

### 🔵 9. **Pipeline Pattern**

* Chaining stages using channels
* Fan-in and fan-out
* Cancellation propagation

### 🔵 10. **Rate Limiting and Throttling**

* Using `time.Ticker`, `time.After`
* Token bucket / leaky bucket strategies
* Use cases: API clients, background jobs

---

## 🔴 Advanced & Performance Topics

### 🔴 11. **Goroutine Leaks**

* How leaks happen
* How to debug leaks (runtime profiles, `pprof`, channel blocking)

### 🔴 12. **Select and Channel Starvation**

* Channel starvation scenarios
* Fairness with select
* Using `reflect.Select` for dynamic cases

### 🔴 13. **Memory Management and Concurrency**

* Escape analysis and goroutine allocation
* Garbage collector behavior under concurrency

### 🔴 14. **Patterns and Anti-Patterns**

* Don’t pass nil channels to `select`
* Avoid unnecessary goroutines
* Avoid shared mutable state

---

## 🔁 Tooling and Debugging

### 🧰 15. **Concurrency Debugging Tools**

* `go test -race` — Data race detector
* `pprof` — Goroutine, CPU, memory profiling
* `trace` — Visual timeline of goroutines

---

## 📚 Optional but Useful

* Custom channel implementations (circular buffers)
* Priority queues using goroutines
* Event-driven concurrency (watchers/listeners)
* Batching and aggregation (e.g., debounce, coalescing)

---

### 📌 Suggested Learning Path

```text
Goroutines
↓
Unbuffered/Buffered Channels
↓
select + timeouts
↓
sync.WaitGroup + Mutex
↓
Context + Pipeline Patterns
↓
Worker Pool, Rate Limiting
↓
Concurrency Debugging + Anti-Patterns
↓
Mastery!
```

---

## 📘 Understanding Channels in Go (with Examples)

This README demonstrates how **Go channels** behave under different conditions, highlighting:

* Unbuffered vs buffered channels
* Blocking behavior and deadlocks
* Synchronization with `sync.WaitGroup`
* Proper channel usage (e.g., who should close it)

---

## 🧪 Unbuffered Channel Basics

### ❌ Deadlock Example 1 — No Receiver

```go
myChan := make(chan int)
myChan <- 76 // ❌ Deadlock: no receiver
```

> **Explanation**: Unbuffered channels block until both sender and receiver are ready. Since no goroutine is reading from the channel, the send blocks forever.

---

### ❌ Deadlock Example 2 — Receiver Too Late

```go
myChan := make(chan int)
myChan <- 76 // ❌ Still blocks

go func() {
    fmt.Println(<-myChan)
}()
```

> **Explanation**: Even though a receiver exists, it is created *after* the blocking send. The program deadlocks because the sender never completes, and the receiver is never scheduled.

---

### ✅ Working Example — Receiver First

```go
myChan := make(chan int)

go func() {
    fmt.Println(<-myChan)
}()

myChan <- 76 // ✅ This works!
```

> **Explanation**: The receiver is already listening when the send occurs, so the send proceeds and the value is received.

---

## 💾 Buffered Channels

### ✅ Buffered Send and Receive

```go
bufferChan := make(chan int, 1)
bufferChan <- 76         // ✅ Does not block: space in buffer
fmt.Println(<-bufferChan) // ✅ Outputs: 76
```

---

### ❌ Buffer Overflow (Deadlock)

```go
bufferChan := make(chan int, 1)
bufferChan <- 76         // OK
bufferChan <- 77         // ❌ Deadlock: buffer is full
```

> **Fix**: To avoid this, read from the buffer before trying to send another value.

---

## 🔁 Range Over a Channel

### ❌ Receiver Blocking on Closed Channel

```go
bufferChan := make(chan int, 1)
bufferChan <- 76

for val := range bufferChan {
    fmt.Println(val)
}
// ❌ Deadlock: the loop never exits unless channel is closed
```

> **Fix**: You must close the channel for `range` to finish.

---

## ✅ Concurrent Example With `sync.WaitGroup`

Here’s a fixed and working version of your full example:

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    bufferChan := make(chan int, 1)
    wg := sync.WaitGroup{}
    wg.Add(2) // We'll wait for both sender and receiver

    // Sender goroutine
    go func() {
        senderWg := sync.WaitGroup{}
        for i := 0; i < 100; i++ {
            senderWg.Add(1)
            go func(val int) {
                defer senderWg.Done()
                time.Sleep(10 * time.Millisecond)
                bufferChan <- val
            }(i)
        }
        senderWg.Wait()
        close(bufferChan) // ✅ Only sender should close
        wg.Done()
    }()

    // Receiver goroutine
    go func() {
        for val := range bufferChan {
            fmt.Println(val)
        }
        wg.Done()
    }()

    wg.Wait()
}
```

---

## ✅ Key Concepts Recap

| Concept             | Explanation                                                           |
| ------------------- | --------------------------------------------------------------------- |
| Unbuffered channels | Block until both sender and receiver are ready.                       |
| Buffered channels   | Allow sending without blocking — up to capacity.                      |
| `range` on channels | Only exits when the channel is **closed**.                            |
| Closing a channel   | Only the **sender** should close a channel.                           |
| `sync.WaitGroup`    | Used to coordinate goroutines and avoid exiting early or deadlocking. |

---

## ❗ Common Misconceptions (Fixed)

| Misconception                             | Correction                                                                        |
| ----------------------------------------- | --------------------------------------------------------------------------------- |
| “Receiver can close the channel”          | ❌ **Only the sender** should close the channel.                                   |
| “Buffered channel won’t block”            | ❌ It still blocks **once full**.                                                  |
| “Send and receive in same function works” | ❌ Not without a goroutine — both will block.                                      |
| `for i := range 100` is valid             | ❌ Use `for i := 0; i < 100; i++` instead.                                         |
| Channels work like slices                 | ❌ Channels are synchronous (unless buffered), and behave differently from slices. |

---

## 🧠 Final Tips for Interviews

* Always mention that **unbuffered channels are blocking**.
* Use `go func() {}` to run either the sender or receiver concurrently.
* Don’t forget to **close channels** when you're done — especially when ranging.
* Explain **why only the sender should close** the channel.
* Use `sync.WaitGroup` to safely wait for goroutines to finish.

---
