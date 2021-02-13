package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

const APIURL = "http://localhost:8080/users"

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
		log.Fatal(err)
	}
}

func main() {
	// Post new User
	body := bytes.NewBuffer([]byte(
		`{"name":"admin", "email":"admin@gmail.com", "password": "admin", "age": 27, "active": true}`))
	resp, err := http.Post(APIURL, "application/json", body)
	//req, err := http.NewRequest("POST", APIURL, body)
	//req.Header.Add("Accept", `Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	//req.Header.Add("Content-Type", `application/json`)
	//resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	PrintResponse(resp)

	// Get all Users
	resp, err = http.Get(APIURL)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	PrintResponse(resp)
}
