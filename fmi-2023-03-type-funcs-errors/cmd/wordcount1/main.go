package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) (counts map[string]int) {
	counts = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counts[word]++
	}
	return
}

func main() {
	sentence := "to go or not to go this is the go question"
	fmt.Printf("%v\n", WordCount(sentence))
}
