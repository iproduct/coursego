package main

import "fmt"

func IndexRune(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}

func main() {
	f := func (s string, r rune) int {
		for i, c := range s {
			if c == r {
				return i
			}
		}
		return -1
	}
	fmt.Printf("%c in %s -> %d\n", 'o', "Google", f("Google", 'o' ))
	fmt.Printf("%c in %s -> %d\n", 'r', "Google", IndexRune("Google", 'r' ))
}
