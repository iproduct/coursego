package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	userPath := "/users/123?sortBy=name"
	u, _ := url.Parse(userPath)
	p := u.Path
	lastSegment := path.Base(p)
	sortBy := u.Query().Get("sortBy")
	fmt.Printf("Path: %#v\n", p)
	fmt.Printf("UserID: %#v\n", lastSegment)

	fmt.Printf("Sorting by: %#v\n", sortBy)
	fmt.Printf("%#v\n", u)
}