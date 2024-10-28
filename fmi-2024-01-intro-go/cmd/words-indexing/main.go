package main

import (
	"fmt"
	"regexp"
)

func WordCount(s string, counts map[string]int) map[string]int {
	//words := strings.Fields(s)
	re := regexp.MustCompile(`\W+`)
	words := re.Split(s, -1)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func main() {
	counts := make(map[string]int)
	var str string
	for {
		n, err := fmt.Scanf("%s", &str)
		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			WordCount(str, counts)
		}
		if n == 0 {
			break
		}
	}
	fmt.Printf("%v\n", counts)
}
