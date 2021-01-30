package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M() // (<nil>, <nil>)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x49c106]
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
