package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	userPath := "/users/123?start=1000&maxresults=50"
	u, _ := url.Parse(userPath)
	p := u.Path
	lastSegment := path.Base(p)
	sortBy := u.Query().Get("start")
	maxResults := u.Query().Get("maxresults")
	fmt.Printf("Path: %#v\n", p)
	fmt.Printf("UserID: %#v\n", lastSegment)

	fmt.Printf("Start by: %#v\n", sortBy)
	fmt.Printf("Page size: %#v\n", maxResults)
	fmt.Printf("%#v\n", u)
}