package main

import "fmt"

func main() {
	a1 := [...]int{1, 2}
	a2 := a1
	a2[0] = 3
	fmt.Printf("%v , %v , %t \n", a1, a2, a1 == a2)
}
