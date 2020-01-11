package main

import "fmt"

func main2() {
	// var a [10]int
	// //a :=  [10]int{1,2,3,4,5,6,7,8,9,10}
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Element: a[%d] = %d\n", i, a[i])
	// }
	// var a1 [100]int
	// var matrix [4][4]float64
	// a2 := [...]int{1, 1, 2, 3, 5}
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(len(a)) // len == cap
	fmt.Println(cap(a))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	a1 := [...]int{1, 2}
	a2 := a1
	a2[0] = 3
	fmt.Printf("%v, %v, %t, %t\n", a1, a2, a1 == a2, &a1 == &a2)

	var a3 [20]int
	firstHalf := a3[:10]
	secondHalf := a3[10:]
	middle := a3[5:15]
	all := a3[:]
	fmt.Printf("%v, %v, %v, %v\n", firstHalf, secondHalf, middle, all)

	var slice []int = []int{2, 3, 5, 7, 11, 13}
	fmt.Println(slice) // [2 3 5 7 11 13]
	reslice := slice[2:5]
	fmt.Println(reslice) // [5 7 11]

}

