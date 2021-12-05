package main

import (
	"fmt"
	"sync"
)

const NUM_GOROUTINES = 100

var n int64 = 0
var mu sync.Mutex

func increment(id int, wg *sync.WaitGroup) {
	for i := 0; i < 10000; i++ {
		mu.Lock()
		n++
		//fmt.Printf("Updating %d -> %d\n", id, i)
		mu.Unlock()
	}
	wg.Done()
}

func main() {
	//runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for i := 0; i < NUM_GOROUTINES; i++ {
		wg.Add(1)
		go increment(i, &wg)
	}
	wg.Wait()
	fmt.Println(n)
}
