package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	for _, word := range strings.FieldsFunc(s, func(r rune) bool { return r == ' ' }) {
		result[word] += 1
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
