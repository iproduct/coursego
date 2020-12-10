package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan rxgo.Item)
	go func() {
		defer wg.Done()
		for i := 0; i < 30; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
	wg.Add(1)
	observable := rxgo.FromChannel(ch)

	// First Observer
	go func() {
		defer wg.Done()
		for item := range observable.Observe() {
			fmt.Println("First observer: ", item.V)
		}
	}()
	wg.Add(1)

	// Second Observer
	go func() {
		defer wg.Done()
		for item := range observable.Observe() {
			fmt.Println("Second observer: ", item.V)
		}
	}()
	wg.Add(1)

	wg.Wait()
}