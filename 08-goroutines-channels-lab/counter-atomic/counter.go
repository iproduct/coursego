package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 = 0
	for i := 0; i < 10000; i++ {
		go func() {
			num := atomic.AddInt64(&counter, 1)
			fmt.Printf("goroutine %d Counter = %d\n", i, num)
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Counter = %d\n", atomic.LoadInt64(&counter))
}
