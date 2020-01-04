package main

import (
	"fmt"
)

func main() {
	counts := make(map[string]int)

	for {
		var str string
		n, _ := fmt.Scanln(&str)
		if n == 0 {
			break
		}

		fmt.Printf("Entered: %s\n", str)
		counts[str]++
	}

	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%-20.20s -> %d\n", line, n)
		}
	}
}
