package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	for i := 0; i < 10000; i++ {
		go func() {
			counter++
			fmt.Printf("goroutine %d Counter = %d\n", i, counter)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("Counter = %d\n", counter)
}
