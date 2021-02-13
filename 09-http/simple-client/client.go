package main

import (
	"os"
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// Issue an HTTP GET request to a simple-server. `09-http.Get` is a
	// convenient shortcut around creating an `09-http.Client`
	// object and calling its `Get` method; it uses the
	// `09-http.DefaultClient` object which has useful default
	// settings.
	var resp *http.Response
	var err error
	if len(os.Args) > 1 {
		resp, err = http.Get(os.Args[1])
	} else {
		resp, err = http.Get("09-http://gobyexample.com")
	}
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
