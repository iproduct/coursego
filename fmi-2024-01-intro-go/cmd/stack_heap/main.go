package main

import "fmt"

var global *int

func f() int {
	var x int
	x = 1
	global = &x
	return 42
}

func g() int {
	y := new(int)
	*y++
	return *y
}

func main() {
	fmt.Println(f())
	fmt.Println(g())
}
