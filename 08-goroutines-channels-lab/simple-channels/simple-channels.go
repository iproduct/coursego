package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ProduceEventsWG() <-chan string {
	words := []string{
		"Hello", "Goroutines", "and", "Channels", "from", "Go", "Language",
	}
	out := make(chan string)
	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			out <- w
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}(word)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ProduceEventsChannel() <-chan string {
	words := []string{
		"Hello", "Goroutines", "and", "Channels", "from", "Go", "Language",
	}
	out := make(chan string)
	done := make(chan struct{})
	for _, word := range words {
		go func(w string) {
			defer func() {
				done <- struct{} {}
			}()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- w
		}(word)
	}
	go func() {
		for range words {
			<-done
		}
		close(done)
		close(out)
	}()
	return out
}

func main() {
	wordsChannel := ProduceEventsChannel()
	for message := range wordsChannel {
		fmt.Println(message)
	}
}
