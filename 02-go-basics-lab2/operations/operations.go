package main

import "fmt"

func main() {
	a := 0b10101010
	b := 0b00001111
	fmt.Printf("%b & \n%08b = \n%08b\n", a, b, a&b)
}
