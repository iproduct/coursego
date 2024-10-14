package main

import "fmt"

func main() {
	for i, c := range "go" {
		fmt.Printf("index: %d, rune: %c\n", i, c)
		fmt.Println(i, string(c))
	}
}
