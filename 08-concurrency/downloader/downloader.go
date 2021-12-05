package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {
	quit := make(chan struct{})
	urlGen := UrlGenerator(5, quit)

	resources := DownloadResources(urlGen, quit)
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

func UrlGenerator(maxNumber int, quit <-chan struct{}) <-chan string {
	rand.Seed(time.Now().Unix())
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < maxNumber; i++ {
			url := fmt.Sprintf("http://example.com/resource/%d/%d", i, rand.Intn(100)*100)
			select {
			case out <- url:
				fmt.Printf("Generating URL: %s\n", url)
			case <-quit:
				return
			}
		}
	}()
	return out
}

type Resource struct {
	size int
	err  error
}

func DownloadResources(urls <-chan string, quit chan struct{}) <-chan Resource {
	resources := make(chan Resource)
	go func() {
		defer close(resources)
		var wg sync.WaitGroup
		for u := range urls {
			if IsClosed(quit) {
				return
			}
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				if IsClosed(quit) {
					fmt.Printf("Quiting goroutine for url: %s\n", url)
					return
				}
				resource := download(url)
				if IsClosed(quit) {
					fmt.Printf("Quiting goroutine for url: %s\n", url)
					return
				}
				resources <- resource
				if resource.err != nil && !IsClosed(quit) { // guard not to close quit channel twice, resulting in panic
					close(quit) // not yet closed - close it now
				}
			}(u)
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

func download(urlStr string) Resource {
	random := rand.Float64()
	fmt.Printf("Downloading URL: %s\n", urlStr)
	switch {
	case random < 0.8:
		url, err := url.Parse(urlStr)
		if err != nil {
			log.Fatal(err)
		}
		size, err := strconv.Atoi(path.Base(url.Path))
		if err != nil {
			return Resource{0, err}
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
	return total, err
}
