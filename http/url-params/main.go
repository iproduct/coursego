package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {
	userPath := "/users/123?sortBy=name"
	u, _ := url.Parse(userPath)
	ps := path.Base(u.Path)
	sortBy := u.Query().Get("sortBy")
	fmt.Printf("UserID: %#v\n", ps)
	fmt.Printf("Sorting by: %#v\n", sortBy)
	fmt.Printf("%#v\n", u)
}