package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://localhost:8080/headers")
	//resp, err := http.Get("http://google.com")
	req, err := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	req.Header.Add("Accept", `Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("Custom-Header", `Custom Value`)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(i+1, ": ", scanner.Text())
	}
}
