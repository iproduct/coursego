package main

import "fmt"

type S struct {
	a int
	b float64
}

func main() {
	s1 := new(S)
	fmt.Printf("%#v", s1)
}
