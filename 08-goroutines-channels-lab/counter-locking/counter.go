package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			counter++
			n := counter
			mu.Unlock()
			fmt.Printf("goroutine %d Counter = %d\n", i, n)
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Counter = %d\n", counter)
}
