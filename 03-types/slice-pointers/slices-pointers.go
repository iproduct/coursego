package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names) // [John Paul George Ringo]

	a := names[1:2]
	fmt.Printf("a: %#v, len = %d, cap = %d\n", a, len(a), cap(a)) // [Paul]
	b := a[2:3]
	fmt.Printf("b: %#v, len = %d, cap = %d\n", b, len(b), cap(b)) // [Ringo]

	b[0] = "XXX"
	fmt.Println(a, b)  // [Paul] [XXX]
	fmt.Println(names) // [John Paul George XXX]
}
