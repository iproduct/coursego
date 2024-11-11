package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://localhost:8080/headers")
	//req, err := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	//req, err := http.NewRequest("GET", "https://www.googleapis.com/books/v1/volumes?q=golang", nil)
	req, err := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Response Status: %v\n", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 40; i++ {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
