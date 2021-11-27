package main

import "fmt"

func process1(x [3]string) int {
	x[0] = "New"
	sum := 0
	for _, v := range x {
		sum += len(v)
	}
	return sum
}

func process2(x *[3]string) int {
	x[0] = "New"
	sum := 0
	for _, v := range x {
		sum += len(v)
	}
	return sum
}

func lengths(x [3]string) [3]int {
	var results [3]int
	for i, v := range x {
		results[i] += len(v)
	}
	return results
}

func main() {
	a := [...]string{"Hello", "Golang", "!"}

	//var a [3]string
	//a[0] = "Hello"
	//a[1] = "Golang"

	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := a
	b[1] = "Java"

	for i, s := range a {
		fmt.Printf("a[%d] : %#s\n", i, s)
	}

	for i, s := range b {
		fmt.Printf("b[%d] : %#s\n", i, s)
	}

	fmt.Println(process2(&a))

	for i, s := range a {
		fmt.Printf("a[%d] : %#s\n", i, s)
	}

	c := lengths(a)
	fmt.Println(c)
}
