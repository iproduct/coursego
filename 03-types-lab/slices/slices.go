package main

import "fmt"

func main() {
	a := make([]int, 5)    // len(a)=5
	printSlice("a", a)     // a len=5 cap=5 [0 0 0 0 0]
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice("b", b)     // b len=0 cap=5 []
	b = b[:cap(b)]         // len(b)=5, cap(b)=5
	printSlice("b", b)     // b len=5 cap=5 [0 0 0 0 0]
	b = b[1:]              // len(b)=4, cap(b)=4
	printSlice("b", b)     // b len=4 cap=4 [0 0 0 0]

	var c []int = nil
	printSlice("c", c)

	var m map[int]string = make(map[int]string)
	m[42] = "The answer"
	fmt.Printf("%#v, %[1]p\n", m)
	//a := make([]int, 5, 10)
	//printSlice("a", a) // a len=5 cap=10 [0 0 0 0 0]
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b) // b len=0 cap=5 []
	//
	//c := b[:2]
	//printSlice("c", c) // c len=2 cap=5 [0 0]
	//
	//d := c[2:4:5]
	//printSlice("d", d) // d len=2 cap=3 [0 0]
	//
	//e := a[2:5:10]
	//printSlice("e", e) // e len=3 cap=8 [0 0 0]
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
