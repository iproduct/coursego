package main

import "fmt"

func main() {
	var a [10]int
	fmt.Println(a, len(a), cap(a))
	b := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(b, len(b), cap(b))
	a = b
	for i := 0; i < len(a); i++ {
		a[i] += 10
	}
	fmt.Println(a, len(a), cap(a))
	fmt.Println(b, len(b), cap(b))

	sb := b[:5]
	fmt.Println("sb = ", sb, len(sb), cap(sb))
	sa := sb[2:7]
	for i := 0; i < len(sa); i++ {
		sa[i] *= sa[i]
	}
	fmt.Println("sa = ", sa[:8], len(sa), cap(sa))
	fmt.Println("sb = ", sb, len(sb), cap(sb))
}
