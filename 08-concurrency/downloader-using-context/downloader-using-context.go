package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	urlGen := UrlGenerator(ctx, 5000)
	resources := DownloadResources(ctx, cancel, urlGen)
	//for res := range resources {
	//	fmt.Printf("Url: %+v\n", res)
	//}
	total, err := CalculateTotalSize(resources)
	if err != nil {
		fmt.Printf("In main() error: %s\n", err)
	} else {
		fmt.Printf("Download success. Totaly downloaded: %d bytes\n", total)
	}
	fmt.Printf("Final number of goroutines: %d\n", runtime.NumGoroutine())
}

func UrlGenerator(ctx context.Context, maxNumber int) <-chan string {
	rand.Seed(time.Now().Unix())
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < maxNumber; i++ {
			url := fmt.Sprintf("http://example.com/resource/%d/%d", i, rand.Intn(100)*100)
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

type Resource struct {
	size int
	err  error
}

func DownloadResources(ctx context.Context, cancel context.CancelFunc, urls <-chan string) <-chan Resource {
	resources := make(chan Resource)
	go func() {
		defer close(resources)
		var wg sync.WaitGroup
		number := 0
		for u := range urls {
			wg.Add(1)
			go func(url string, n int) {
				defer wg.Done()
				resource := download(ctx, url)
				select {
				case resources <- resource:
					if resource.err != nil {
						fmt.Printf("Cancelling the whole pipeline for url: %s\n", url)
						cancel() //can be called safely from multiple goroutines - after first call it does nothing
					}
				case <-ctx.Done():
					fmt.Printf("Canceling Downloader %d\n", n)
					fmt.Printf("Quiting goroutine for url: %s\n", url)
				}
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}(u, number)
			number++
		}
		wg.Wait()
	}()
	return resources
}

func IsClosed(ch <-chan struct{}) bool {
	select {
	case <-ch:
		return true //already closed - return true
	default:
		return false // not yet closed - return false
	}
}

func download(ctx context.Context, urlStr string) Resource {
	random := rand.Float64()
	fmt.Printf("Downloading URL: %s\n", urlStr)
	switch {
	case random >= 0:
		url, err := url.Parse(urlStr)
		if err != nil {
			return Resource{0, err}
		}
		size, err := strconv.Atoi(path.Base(url.Path))
		if err != nil {
			return Resource{0, err}
		}
		if ctx.Err() != nil {
			return Resource{0, ctx.Err()}
		}
		fmt.Printf("Successfully downloaded resource %s -> size: %d\n", urlStr, size)
		return Resource{size, nil}
	default:
		fmt.Printf("Error downloading resource: %s\n", urlStr)
		return Resource{0, fmt.Errorf("Error downloading resource: %s", urlStr)}
	}
}

func CalculateTotalSize(resources <-chan Resource) (int, error) {
	total := 0
	var err error = nil
	for resource := range resources {
		total += resource.size
		err = resource.err
	}
	fmt.Println("Exiting Calculation")
	return total, err
}
