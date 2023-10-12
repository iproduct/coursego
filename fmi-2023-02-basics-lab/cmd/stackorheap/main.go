package main

import "fmt"

var global *int

func f() {
	var x int
	x = 42
	global = &x
}

func g() {
	var y = 42
	fmt.Printf("%d", y)
}

func main() {
	f()
	fmt.Printf("%d", *global)

}
