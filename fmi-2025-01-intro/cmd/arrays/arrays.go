package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//var a [10]int
	//a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%d ", a[i])
	//}
	//var a1 [100]int
	//fmt.Printf("%d ", a1)
	//fmt.Println()
	//var matrix [4][4]float64
	//fmt.Println(matrix)
	////a2 := [...]int{1, 1, 2, 3, 5}
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//a_slice := a[:]
	//a_slice_new := append(a_slice, "from")
	//a_slice_two := append(a_slice_new, "Golang")
	//
	//fmt.Printf("%v -> %p - %v - %v - %v\n", a_slice, &a_slice, unsafe.SliceData(a_slice), len(a_slice), cap(a_slice))
	//fmt.Printf("%v -> %p - %v - %v - %v\n", a_slice_new, &a_slice_new, unsafe.SliceData(a_slice_new), len(a_slice_new), cap(a_slice_new))
	//fmt.Printf("%v -> %p - %v - %v - %v\n", a_slice_two, &a_slice_two, unsafe.SliceData(a_slice_two), len(a_slice_two), cap(a_slice_two))
	// fmt.Println(a[0], a[1])
	// fmt.Println(len(a)) // len == cap
	// fmt.Println(cap(a))

	// primes := [6]int{2, 3, 5, 7, 11, 13}
	// fmt.Println(primes)

	//a5 := [...]int{1, 2, 42}
	//a6 := a5
	////a6[2] = 108
	//fmt.Printf("%v -> %p\n", a5, &a5)
	//fmt.Printf("%v -> %p\n", a6, &a6)
	////for _, e := range a6 {
	////	fmt.Printf("%v ", e)
	////}
	//////a2[0] = 3
	//fmt.Printf("%v, %v, %t, %t\n", a5, a6, a5 == a6, &a5 == &a6)

	var a3 [20]int
	firstHalf := a3[:10]
	secondHalf := a3[10:]
	middle := a3[5:15]
	all := middle[5:]
	all2 := all[:]
	appended := append(secondHalf, 17)
	appended[0] = 42
	fmt.Printf("firstHalf: %v, len: %d, cap: %d\n", firstHalf, len(firstHalf), cap(firstHalf))
	fmt.Printf("secondHalf: %v, len: %d, cap: %d -> %p\n", secondHalf, len(secondHalf), cap(secondHalf), unsafe.SliceData(secondHalf))
	fmt.Printf("middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))
	fmt.Printf("all: %v, len: %d, cap: %d\n", all, len(all), cap(all))
	fmt.Printf("all2: %v, len: %d, cap: %d\n", all2, len(all2), cap(all2))
	fmt.Printf("%v, len: %d, cap: %d -> %p\n", appended, len(appended), cap(appended), unsafe.SliceData(appended))
	fmt.Printf("%v, len: %d, cap: %d\n", a3, len(a3), cap(a3))
	//  %v, %v, %v\n", firstHalf, secondHalf, middle, all)

	var slice []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(slice) // [2 3 5 7 11 13]
	reslice := slice[2:4:5]
	reslice2 := reslice[1:2:3]
	fmt.Println(reslice, len(reslice), cap(reslice))    // [5 7 11]
	fmt.Println(reslice2, len(reslice2), cap(reslice2)) //

}
