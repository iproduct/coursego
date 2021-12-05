package main

import (
	"fmt"
	"sync"
	"time"
)

// Lock Contention is the process when a process or thread tries to acquire a mutex-atomicint that is held by another
// process or thread, thus causing it to wait longer than it needs to
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		fmt.Println("go routine 1 trying to acquire mutex", time.Now())
		mu.Lock()
		fmt.Println("go routine 1 locking", time.Now())
		time.Sleep(3 * time.Second)
		fmt.Println("go routine 1 releasing mutex after 3s:", time.Now())
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		fmt.Println("go routine 2 trying to acquire mutex:", time.Now())
		mu.Lock()
		fmt.Println("go routine 2 acquired mutex after 3s:", time.Now())
		time.Sleep(1 * time.Second)
		fmt.Println("go routine 2 releasing mutex after 1s:", time.Now())
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
