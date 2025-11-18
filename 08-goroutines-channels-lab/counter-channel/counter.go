package main

import (
	"fmt"
	"time"
)

func main() {
	counterChannel := make(chan int)
	go func() {
		for counter := 0; ; counter++ {
			counterChannel <- counter
		}
	}()
	for i := 0; i < 10000; i++ {
		go func(k int) {
			n := <-counterChannel
			fmt.Printf("goroutine %d Counter = %d\n", k, n)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("Counter = %d\n", <-counterChannel)
}
