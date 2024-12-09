package main

import (
	"fmt"
	"sync"
)

var count int64

func counter(ids chan<- int64) {
	var wg sync.WaitGroup
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++     // Wrong!
			id := count // Defensive copying
			ids <- id
		}()
	}
	wg.Wait()
	close(ids)
}

func main() {
	ids := make(chan int64)
	idsSet := make(map[int64]struct{})
	go counter(ids)
	for id := range ids {
		idsSet[id] = struct{}{}
		//fmt.Printf("%d, ", id)
	}
	fmt.Printf("\nCounter: %d, Unique IDs count: %d\n", count, len(idsSet))
}
