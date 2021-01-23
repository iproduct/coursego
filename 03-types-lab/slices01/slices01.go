package main

import "fmt"

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		fmt.Printf("Element: a[%d] = %d\n", i, a[i])
	}
	s := a[2:4]
	fmt.Printf("%#v: len: %d, cap:%d", s, len(s), cap(s))

	s1 := append(s, 10, 11, 12, 13, 14, 15, 16, 17, 18)
	fmt.Printf("%#v: len: %d, cap:%d, same:%t", s1, len(s1), cap(s1), &s[0] == &s1[0])
}
