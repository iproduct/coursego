package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func WordCount(s string, counts map[string]int) map[string]int {
	//words := strings.Fields(s)
	re := regexp.MustCompile(`\W+`)
	words := re.Split(s, -1)
	fmt.Println(s, words)
	for _, word := range words {
		if word != "" {
			counts[word]++
		}
	}
	return counts
}

func main() {
	counts := make(map[string]int)
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		//var str string
		//n, err := fmt.Scanf("%s", &str)
		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		str = strings.TrimSpace(str)
		if str == "" {
			break
		} else {
			WordCount(str, counts)
		}
	}
	fmt.Printf("%v\n", counts)
}
