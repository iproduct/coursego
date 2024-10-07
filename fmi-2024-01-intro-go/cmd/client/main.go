package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://googleapis.com/books/v1/volumes?q=golang")
	req, err := http.NewRequest("GET", "https://www.googleapis.com/books/v1/volumes?q=golang", nil)
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
	for i := 0; scanner.Scan() && i < 20; i++ {
		fmt.Println(scanner.Text())
	}

}
