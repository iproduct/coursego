package main

import "fmt"

func swapVal(x, y string) (string, string) {
	return y, x
}
func swapRef(x, y *string) {
	*x, *y = *y, *x
}
func main() {
	a, b := swapVal("hello", "world")
	fmt.Println(a, b)
	swapRef(&a, &b)
	fmt.Println(a, b)
}
