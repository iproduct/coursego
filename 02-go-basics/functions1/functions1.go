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

func g() {
	//y := new(interface{a() int; b() int})
	y := new([5]int)
	fmt.Printf("%#v %[1]T\n", *y)
	//z := y[:2]
	fmt.Println(len(y), cap(y))
	//*y = 1
}


func main() {
	a := &[]int{1,2,3}
	fmt.Printf("%#v %[1]T\n", a)
	g()
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
