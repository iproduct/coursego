package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The call to Wait does the following under the hood
// 1. Calls Unlock() on the condition Locker
// 2. Notifies the list wait
// 3. Calls Lock() on the condition Locker

// The Cond type besides the Locker also has access to 2 important methods:
// 1. Signal - wakes up 1 go routine waiting on a condition (rendezvous point)
// 2. Broadcast - wakes up all go routines waiting on a condition (rendezvous point)

func condition(n *uint64) bool {
	return atomic.LoadUint64(n)%2 == 0
}

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	var n uint64 = 0

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			cond.L.Lock()
			for !condition(&n) {
				cond.Wait()
			}
			// ... make use of condition ...
			atomic.AddUint64(&n, 1)
			fmt.Println("go routine 1")
			cond.L.Unlock()
			cond.Signal()
			//cond.Broadcast()

		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			cond.L.Lock()
			for condition(&n) {
				cond.Wait()
			}
			// ... make use of condition ...
			atomic.AddUint64(&n, 1)
			fmt.Println("go routine 2")
			cond.L.Unlock()
			cond.Signal()
			//cond.Broadcast()

		}
	}()

	wg.Wait()
}
