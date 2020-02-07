package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

const APIUrl = "http://localhost:8088/users"

func PrintResponse(resp *http.Response) {
	// Print the HTTP response status.
	fmt.Println("\nResponse status:", resp.Status)
	fmt.Println("Response headers:", resp.Header)

	// Print the the response body.
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	var resp *http.Response
	var err error

	// Post new User
	resp, err = http.Post(APIUrl, "text/json", bytes.NewBuffer([]byte(`{"name":"admin", "email":"admin@gmail.com"}`)))
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	PrintResponse(resp)

	// Get all Users
	resp, err = http.Get(APIUrl)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	PrintResponse(resp)
}
