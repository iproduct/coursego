package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[1:2]
	b := a[2:3]
	fmt.Println(a, b) // [Paul] [Ringo]

	b[0] = "XXX"
	fmt.Println(a, b)  // [Paul] [XXX]
	fmt.Println(names) // [John Paul George XXX]
}
