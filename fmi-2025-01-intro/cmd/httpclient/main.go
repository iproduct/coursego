package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	//resp, err := http.Get("http://localhost:8080/headers")
	req, err := http.NewRequest("GET", "http://localhost:8080/headers", nil)
	req.Header.Add("Accept", "text/html,application/json")
	req.Header.Add("Custom-Header", "Custom Value")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
