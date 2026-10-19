package main

import "fmt"

func main() {
	var s []int
	//s = []int{}
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
	}
	s = append(s, 42)
	fmt.Println(s, len(s), cap(s)) // [] 0 0
}
