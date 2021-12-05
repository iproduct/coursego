package main

import (
	"fmt"
	"sync"
)

const NUM_GOROUTINES = 100

var n int64 = 0

func inrement(id int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		n++
		//fmt.Printf("Goroutine %v -> %v\n", id, n)
	}
	wg.Done()
}

func main() {
	//runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	for i := 0; i < NUM_GOROUTINES; i++ {
		wg.Add(1)
		go inrement(i, &wg)
	}
	wg.Wait()
	fmt.Println(n)
}
