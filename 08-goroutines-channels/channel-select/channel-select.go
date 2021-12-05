package main

import (
	"fmt"
	"runtime"
	"time"
)

func fibonacci(quit <-chan struct{}) <-chan int {
	fibChannel := make(chan int)
	go func() {
		defer close(fibChannel)
		fmt.Println("Generating fibonacci numbers ...")
		a, b := 0, 1
		for {

			select {
			case fibChannel <- a:
				fmt.Printf("a = %d\n", a)
				a, b = b, a+b
			case <-quit:
				fmt.Println("Canceling fibonacci generation.")
				return
			default:
				fmt.Println("No activity - sleeping for 100 ms ...")
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return fibChannel
}

func main() {

	quitChannel := make(chan struct{})
	fibChannel := fibonacci(quitChannel)
	fmt.Println("Fibonacci consumer goroutine started ...")
	for i := 0; i < 10; i++ {
		value := <-fibChannel
		fmt.Printf("Consuming Fibonacci [%d] = %d\n", i, value)
	}
	quitChannel <- struct{}{}
	close(quitChannel)
	fmt.Println("Starting fibonacci generator ...")
	fmt.Printf("Final number of goroutines: %d\n", runtime.NumGoroutine())
	// Make a copy of MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)
	fmt.Printf("Memstats: %#v\n", m0)
}
