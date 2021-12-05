package main

import (
	"context"
	"fmt"
	concurrentset "github.com/iproduct/coursego/06-coroutines-channels/concurrent-set"
	"github.com/iproduct/coursego/06-coroutines-channels/semaphor"
	"math/rand"
	"runtime"
	"sync/atomic"
	"time"
)

// Fake a long and difficult work.
func DoWork(url string,
	ctx context.Context,
	visited *concurrentset.ConcurrentHashSet,
	numUrls *uint64,
	urls chan<- string,
	jobs semaphor.Semaphor,
	) {
	fmt.Println("doing", url)
	time.Sleep(500 * time.Millisecond)
	fmt.Println("finished", url)
	defer jobs.Release() // release the token

	for i := 0; i < 10; i++ {
		newUrl := fmt.Sprintf("%s/%d", url, i)
		num := atomic.AddUint64(numUrls, 1)
		fmt.Printf("sending new URL %d: %s\n", num, newUrl)
		if visited.IsMember(newUrl) || atomic.LoadUint64(numUrls) >= MAX_URLS {
			continue
		}
		select {
		case urls <- newUrl:
		case <-ctx.Done():
			return
		}
	}
}

const MAX_URLS = 300

func main() {
	// concurrentJobs is a buffered channel implemting semaphore that blocks
	// if more than 20 goroutines are started at once
	var concurrentJobs = semaphor.New(10)
	numUrls := uint64(0)
	visited := concurrentset.New()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	urls := make(chan string, 100)
	urls <- "http://urls/1"
	for url := range urls {
		if atomic.LoadUint64(&numUrls) >= MAX_URLS {
			cancel()
			break
		}
		concurrentJobs.Acquire() // acquire a  token
		go DoWork(url, ctx, visited, &numUrls, urls, concurrentJobs)
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
