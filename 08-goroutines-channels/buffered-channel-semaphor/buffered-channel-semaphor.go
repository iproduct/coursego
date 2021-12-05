package main

import (
	"context"
	"fmt"
	concurrentset "github.com/iproduct/coursego/06-coroutines-channels/concurrent-set"
	"github.com/iproduct/coursego/06-coroutines-channels/semaphor"
	"math/rand"
	"runtime"
	"time"
)

// Fake a long and difficult work.
func DoWork(url string, jobs semaphor.Semaphor) {
	fmt.Println("doing", url)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finished", url)
	jobs.Release() // release the token
}

const MAX_URLS = 300

func main() {
	// concurrentJobs is a buffered channel implemting semaphore that blocks
	// if more than 20 goroutines are started at once
	var concurrentJobs = semaphor.New(10000)
	visited := concurrentset.New()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	for url := range UrlGenerator(ctx, MAX_URLS, visited) {
		concurrentJobs.Acquire() // acquire a  token
		go DoWork(url, concurrentJobs)
		fmt.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
	}
}

func UrlGenerator(ctx context.Context, maxNumber int, visited *concurrentset.ConcurrentHashSet) <-chan string {
	rand.Seed(time.Now().Unix())
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < maxNumber; i++ {
			url := fmt.Sprintf("http://example.com/resource/%d", rand.Intn(500))
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
