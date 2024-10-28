package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func WordCount(f *os.File) map[string]int {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	split := regexp.MustCompile(`\W+`)
	letter := regexp.MustCompile(`.*[a-zA-Z]{2,}.*`)
	for input.Scan() {
		words := split.Split(input.Text(), -1)
		for _, word := range words {
			if letter.MatchString(word) {
				counts[word]++
			}
		}
	}
	if err := input.Err(); err != nil {
		log.Default().Printf("Error in WordCount: %s", err)
	}
	return counts
}

func main() {
	fmt.Printf("Path: %s\n", os.Args[0])
	files := os.Args[1:]
	fmt.Println("files:", files)
	if len(files) == 0 {
		fmt.Printf("Usage: %s FILE1 FILE2 ...\n", os.Args[0])
	} else {
		for _, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
			defer f.Close()
			counts := WordCount(f)
			fmt.Printf("%v\n", counts)
		}
	}
}
