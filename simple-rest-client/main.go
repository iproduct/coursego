package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

const APIUrl = "http://localhost:8080/users"

func main() {
	var resp *http.Response
	var err error
	resp, err = http.Post(APIUrl, "application/json", bytes.NewBuffer([]byte(`{"name":"admin", "email":"admin@gmail.com"}`)))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 5 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
