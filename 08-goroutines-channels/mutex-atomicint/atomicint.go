package main

import (
	"fmt"
	"sync"
)

// AtomicInt is a concurrent data structure that holds an int.
// Its zero value is 0.
type AtomicInt struct {
	mu sync.Mutex // A mutex-atomicint than can be held by one goroutine at a time.
	n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
	a.mu.Lock() // Wait for the mutex-atomicint to be free and then take it.
	a.n += n
	a.mu.Unlock() // Release the mutex-atomicint.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

func main() {
	wait := make(chan struct{})
	var n AtomicInt
	go func() {
		n.Add(1) // one access
		close(wait)
	}()
	n.Add(1) // another concurrent access
	<-wait
	fmt.Println(n.Value()) // 2
}
