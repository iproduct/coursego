package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("%#v, len = %d, cap = %d\n", q, len(q), cap(q)) // [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // [true false true true false true]

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) // [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}
