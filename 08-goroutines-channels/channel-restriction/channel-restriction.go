package main

import (
	"fmt"
	"math/rand"
)

func randomFeed(count, maxVal int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- rand.Intn(maxVal)
		}
	}()
	return ch
}

func main() {
	intVals := randomFeed(5, 1000)
	for value := range intVals {
		fmt.Println(value)
	}
}
