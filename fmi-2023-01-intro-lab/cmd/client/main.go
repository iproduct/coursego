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
	req, err := http.NewRequest("POST", "http://localhost:8080/headers", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	req.Header.Add("Accept", "Accept: text/html,*/*;q=0.8")
	req.Header.Add("Custom-Header", "Custom value")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 15; i++ {
		fmt.Println(scanner.Text())
	}
}
