package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // [] 0 0
	if s == nil {
		fmt.Println("nil!") // nil!
	}
	s = append(s, 42, 43, 44)
	s = append(s, 45)
	s = append(s, 47, 48, 49)
	fmt.Println(s, len(s), cap(s))
}
