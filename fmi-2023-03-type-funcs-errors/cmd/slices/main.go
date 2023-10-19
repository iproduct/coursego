package main

import "fmt"

func main() {
	var a3 [20]int
	firstHalf := a3[:10]
	secondHalf := a3[10:]
	middle := a3[5:15]
	all := a3
	fmt.Printf("%v , %v , %v , %v\n ", firstHalf, secondHalf, middle, all)
	fmt.Printf("firstHalf=%v len=%d, cap=%d\n", firstHalf, len(firstHalf), cap(firstHalf))
	result := append(firstHalf, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("result=%v len=%d, cap=%d\n", result, len(result), cap(result))

}
