package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const NUM_GOROUTINES = 100

var n int64 = 0

func increment(id int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&n, 1)
		//n++
		//fmt.Printf("Goroutine %v -> %v\n", id, n)
	}
	wg.Done()
}

func main() {
	//runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	for i := 0; i < NUM_GOROUTINES; i++ {
		wg.Add(1)
		go increment(i, &wg)
	}
	wg.Wait()
	fmt.Println(n)
}
