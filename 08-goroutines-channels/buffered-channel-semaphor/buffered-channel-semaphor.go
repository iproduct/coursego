package main

import (
	"context"
	"fmt"
	concurrentset "github.com/iproduct/coursego/06-coroutines-channels/concurrent-set"
	"github.com/iproduct/coursego/06-coroutines-channels/semaphor"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Fake a long and difficult work.
func DoWork(url string, jobs semaphor.Semaphor, wg *sync.WaitGroup) {
	fmt.Println("doing", url)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finished", url, ", goroutines: ", runtime.NumGoroutine())
	wg.Done()
	jobs.Release() // release the token
}

const MAX_URLS = 300

func main() {
	// concurrentJobs is a buffered channel implemting semaphore that blocks
	// if more than 20 goroutines are started at once
	var concurrentJobs = semaphor.New(10)
	var wg sync.WaitGroup
	visited := concurrentset.New()
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	ctx, cancel := context.WithCancel(ctx)
	go time.AfterFunc(5*time.Second, cancel)
	for url := range UrlGenerator(ctx, MAX_URLS, visited) {
		concurrentJobs.Acquire() // acquire a  token
		wg.Add(1)
		go DoWork(url, concurrentJobs, &wg)
		fmt.Printf("Current number of goroutines: %d\n", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("Final number of goroutines: %d\n", runtime.NumGoroutine())
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
			time.Sleep(20 * time.Millisecond)
		}
	}()
	return out
}
