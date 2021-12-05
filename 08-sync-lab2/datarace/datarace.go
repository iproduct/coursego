package main

import (
	"fmt"
)

const NUM_GOROUTINES = 100

var n int64 = 0

func increment(id int, complete chan<- struct{}) {
	for i := 0; i < 10000; i++ {
		n++
		//fmt.Printf("Goroutine %v -> %v\n", id, n)
	}
	complete <- struct{}{}
}

func main() {
	//runtime.GOMAXPROCS(2)
	complete := make(chan struct{})
	for i := 0; i < NUM_GOROUTINES; i++ {
		go increment(i, complete)
	}
	for i := 0; i < NUM_GOROUTINES; i++ {
		<-complete
	}
	fmt.Println(n)
}
