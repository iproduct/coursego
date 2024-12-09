package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"sync"
)

//type httpPkg struct{}

//func (httpPkg) Get(url string) {
//	http.Client
//}

//var myhttp httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"https://play.golang.org/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			if res.StatusCode > 299 {
				log.Printf("Response failed with status code: %d and\nbody: \n", res.StatusCode)
			}
			//body, err := io.ReadAll(res.Body)
			//res.Body.Close()
			body := bufio.NewScanner(res.Body)
			lineNum := 0
			for body.Scan() && lineNum < 10 {
				fmt.Println(body.Text()) // Println will add back the final '\n'
				lineNum++
			}
			fmt.Println("\n---------------------------------------------------------------------------------------\n")
			if err := body.Err(); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
