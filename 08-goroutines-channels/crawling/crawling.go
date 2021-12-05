package main

import (
"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Res struct {
	url string
	body string
	found int // Number of new urls found
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan Res, errs chan error, visited map[string]bool) {
	body, urls, err := fetcher.Fetch(url)
	visited[url] = true
	if err != nil {
		errs <- err
		return
	}

	newUrls := 0
	if depth > 1 {
		for _, u := range urls {
			if !visited[u] {
				newUrls++
				go Crawl(u, depth-1, fetcher, ch, errs, visited)
			}
		}
	}

	// Send the result along with number of urls to be fetched
	ch <- Res{url, body, newUrls}

	return
}

func main() {
	ch := make(chan Res)
	errs := make(chan error)
	visited := map[string]bool{}
	go Crawl("http://golang.org/", 4, fetcher, ch, errs, visited)
	tocollect := 1
	for n := 0; n < tocollect; n++ {
		select {
		case s := <-ch:
			fmt.Printf("found: %s %q\n", s.url, s.body)
			tocollect += s.found
		case e := <-errs:
			fmt.Println(e)
		}
	}

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
