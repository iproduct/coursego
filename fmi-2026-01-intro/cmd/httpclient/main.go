package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	//resp, err := http.Get("http://localhost:8001/headers")
	bodyJson, err := json.Marshal(map[string]string{
		"name": "Trayan",
		"age":  "25",
	})
	//body := bytes.NewBuffer(bodyJson)
	body := bytes.NewReader(bodyJson)
	req, err := http.NewRequest("POST", "http://localhost:8001/headers", body)

	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "text/html,application/json")
	req.Header.Add("Custom-Header", "Custom value")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	fmt.Println("response:")
	for i := 0; scanner.Scan(); i++ {
		fmt.Printf("%d: %s\n", i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
