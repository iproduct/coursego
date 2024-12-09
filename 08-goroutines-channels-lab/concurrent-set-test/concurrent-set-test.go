package main

import (
	"08-goroutines-channels-lab/concurrent-set"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	visited := concurrentset.New()
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	urlGen := UrlGenerator(ctx, 50, visited)
	for url := range urlGen {
		fmt.Printf("URL: %v\n", url)
	}
}

func UrlGenerator(ctx context.Context, maxNumber int, visited *concurrentset.ConcurrentHashSet) <-chan string {
	rand.Seed(time.Now().Unix())
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < maxNumber; i++ {
			url := fmt.Sprintf("http://example.com/resource/%d", rand.Intn(5))
			if visited.IsMember(url) {
				continue
			}
			visited.Add(url)
			select {
			case out <- url:
				fmt.Printf("Generating URL: %s\n", url)
			case <-ctx.Done():
				fmt.Println("Canceling UrlGenerator")
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}
