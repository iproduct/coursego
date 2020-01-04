package main

import (
	"fmt"
)

func printf(format string, args ...interface{}) (int, error) {
	_, err := fmt.Printf(format, args...)
	return len(args), err
}

func main() {
	argsLen, err := printf("%v, %v\n", "abcd", 15)
	if err == nil {
		printf("Number args: %d\n", argsLen)
	} else {
		fmt.Printf("Error: %v\n", err)
	}

	// closure demo
	count := 0
	inc := func() int {
		count++
		return count
	}

	incBy := func(n int) int {
		count += n
		return count
	}

	printf("%d\n", inc())
	printf("%d\n", incBy(10))

}
