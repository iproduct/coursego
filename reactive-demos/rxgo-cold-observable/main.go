package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	observable := rxgo.Defer([]rxgo.Producer{func(_ context.Context, ch chan<- rxgo.Item) {
		for i := 0; i < 3; i++ {
			ch <- rxgo.Of(i)
		}
	}})

	// First Observer
	for item := range observable.Observe() {
		fmt.Println("First observer: ", item.V)
	}

	// Second Observer
	for item := range observable.Observe() {
		fmt.Println("Second observer: ", item.V)
	}
}
