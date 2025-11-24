package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func StringProducer(msg string, n int) <-chan string {
	ch := make(chan string)
	go func(msg string, n int) {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}(msg, n)
	return ch
}

func fanIn(inputs ...<-chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(inputs))
	for _, input := range inputs {
		go func(in <-chan string) {
			defer wg.Done()
			for msg := range in {
				out <- msg
			}
		}(input)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func main() {
	p1 := StringProducer("P1", 20)
	p2 := StringProducer("P2", 20)
	p3 := StringProducer("P3", 20)
	p4 := StringProducer("P4", 20)
	out := fanIn(p1, p2, p3, p4)
	for msg := range out {
		fmt.Println(msg)
	}
}
