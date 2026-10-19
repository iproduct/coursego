package main

import (
	"fmt"
)

var s string = "Hello Go World! 你好世界"

func main() {
	//var s string = "Hello Go World! 你好世界"
	s := "Hello Go World! 你好世界"
	//fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c -> %d\n", s[i], i)
	}

	for i, c := range s {
		fmt.Printf("%c -> %d\n", c, i)
	}

	for n := range 10 {
		fmt.Printf("%v\n", n)
	}

	data := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Array len=%v, cap=%v\n", len(data), cap(data))
	slice := data[:]
	slice = append(slice, 8, 9, 10)
	fmt.Printf("Slice len=%v, cap=%v\n", len(slice), cap(slice))
	for _, n := range slice {
		fmt.Printf("%v\n", n)
	}
	//r := []rune(s)
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("%c -> %d\n", r[i], i)
	//}
	//
	//goquote := quote.Go()
	//fmt.Println(goquote)
	//fmt.Println(stringutil.Reverse(goquote))
	//fmt.Println(stringutil.Reverse(s))

}
