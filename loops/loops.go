package main

import (
	"fmt"
)

func main() {
	n, sum := 1, 0
	for n > 0 {
		fmt.Printf("Enter a number: ")
		fmt.Scanln(&n)
		fmt.Printf("n = %d\n", n)
		sum += n
	}
	fmt.Printf("SUM = %d\n\n", sum)

	// using break
	sum = 0
	for {
		fmt.Printf("Enter a number: ")
		fmt.Scanf("%d\n", &n)
		fmt.Printf("n = %d\n", n)
		if n <= 0 {
			break
		}
		sum += n
	}
	fmt.Printf("SUM = %d\n\n", sum)
}
