package wordcount

import "strings"

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		counts[word]++
	}
	return counts
}
