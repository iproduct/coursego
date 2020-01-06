package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if cont := len(files); cont != 0 {
		for _, fname := range files {
			file, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error %v\n", err)
				continue
			}
			countLines(file, counts)
			file.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}

	// for {
	// 	var str string
	// 	n, _ := fmt.Scanln(&str)
	// 	if n == 0 {
	// 		break
	// 	}
	// 	fmt.Printf("Entered: %s\n", str)
	// 	counts[str]++
	// }

	for key, val := range counts {
		fmt.Printf("%-20.20s -> %5d\n", key, val)
	}
}
