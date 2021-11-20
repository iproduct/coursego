package main

import (
	"fmt"
	"math/rand"
)
const size = 5;

func main() {
	sum := 0
	a := [size]int{1, 2, 3, 4, 5}
	b := a
	for  i := 1; i < 5; i++ {
		a[i] = rand.Intn(100)
	}
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)
	fmt.Printf("a == b: %v, &a == &b: %v, &a %p, &b %p, \n", a == b, &a == &b, &a, &b)

	// 1)
	for i := 1; i < size; i++ {
		sum += a[i]
	}
	fmt.Printf("1) Sum = %d, Avg = %g\n", sum, float64(sum)/size)

	// 2) while
	i, sum := 0, 0
	for i < size {
		sum += a[i]
		i++
	}
	fmt.Printf("2) Sum = %d, Avg = %g\n", sum, float64(sum)/size)

	// 3) while true
	i, sum = 0, 0
	for {
		sum += a[i]
		i++
		if i >= size {
			break
		}
	}
	fmt.Printf("3) Sum = %d, Avg = %g\n", sum, float64(sum)/size)

	// 4)
	strings := [...]string{"hello", "golang", "world"}
	for i, s := range strings {
		fmt.Printf("4) %d -> %s\n", i, s)
	}

	// 5)
	john := map[string]string{"name": "john", "email": "john@abv.bg", "age": "32"}
	for k, v := range john {
		fmt.Printf("5) %s -> %s\n", k, v)
	}

	// 6)
	persons := []map[string]string{
		{"name": "john", "email": "john@abv.bg", "age": "32"},
		{"name": "hristo", "email": "hristo@yahoo.com", "age": "25"},
		{"name": "tsvetelina", "email": "tsveti@gmail.com", "age": "28"},
	}
	for i, person := range persons {
		fmt.Printf("6) %d: {\n", i)
		for k, v := range person {
			fmt.Printf("       %s: %s,\n", k, v)
		}
		fmt.Println("      }")
	}

	// 7)
	for i := 'a'; i < 'd'; i++ {
		fmt.Println(string(i))
	}


}
