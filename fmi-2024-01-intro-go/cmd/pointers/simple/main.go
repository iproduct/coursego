package main

import "fmt"

func main() {
	n := 1
	p := &n
	var p2 *int
	fmt.Printf("%p, %v\n", p, *p)
	fmt.Printf("%p, %v\n", p2, p2)
}
