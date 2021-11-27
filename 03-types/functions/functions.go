package main

import (
	"fmt"
)

func printf(format string, args ...interface{}) (int, error) {
	_, err := fmt.Printf(format, args...)
	return len(args), err
}

// closure demo
func NewPersonCounter() (func() int, func(n int) int) {
	count := 0
	inc := func() int {
		count++
		return count
	}
	incBy := func(n int) int {
		count += n
		return count
	}
	return inc, incBy
}

func main() {
	argsLen, err := printf("%v, %v\n", "abcd", 15)
	if err == nil {
		printf("Number args: %d\n", argsLen)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	pcInc, pcIncBy := NewPersonCounter()

	fmt.Println(pcInc())
	fmt.Println(pcIncBy(10))
	fmt.Println(pcInc())
	fmt.Println(pcIncBy(10))
	fmt.Println(pcInc())
	fmt.Println(pcIncBy(10))
	fmt.Println(pcInc())
	fmt.Println(pcIncBy(10))
}
