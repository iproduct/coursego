package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var value = 42
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(value)
		}()
	}
	wg.Wait()
}
