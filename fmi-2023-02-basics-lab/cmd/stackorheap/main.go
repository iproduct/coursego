package main

import "fmt"

var global *int

func f() {
	var x int
	x = 42
	global = &x
}

func g() {
	y := new(int)
	*y = 15
	//global = y
}

func main() {
	f()
	fmt.Printf("%d", *global)
	g()
}
