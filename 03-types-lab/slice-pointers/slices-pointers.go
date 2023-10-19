package main

import "fmt"

func main() {
	names := [...]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[1:2]
	b := a[2:3]
	fmt.Println(a, len(a), cap(a)) // [Paul]
	fmt.Println(b, len(b), cap(b)) // [Ringo]

	b[0] = "XXX"
	fmt.Println(a, b)  // [Paul] [XXX]
	fmt.Println(names) // [John Paul George XXX]
}
