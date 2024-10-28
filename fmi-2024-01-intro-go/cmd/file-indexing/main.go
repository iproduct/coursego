package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

var stopwordsList = []string{`ourselves`, `hers`, `between`, `yourself`, `but`, `again`, `there`, `about`, `once`,
	`during`, `out`, `very`, `having`, `with`, `they`, `own`, `an`, `be`, `some`, `for`, `do`, `its`, `yours`,
	`such`, `into`, `of`, `most`, `itself`, `other`, `off`, `is`, `s`, `am`, `or`, `who`, `as`, `from`, `him`,
	`each`, `the`, `themselves`, `until`, `below`, `are`, `we`, `these`, `your`, `his`, `through`, `don`, `nor`,
	`me`, `were`, `her`, `more`, `himself`, `this`, `down`, `should`, `our`, `their`, `while`, `above`, `both`,
	`up`, `to`, `ours`, `had`, `she`, `all`, `no`, `when`, `at`, `any`, `before`, `them`, `same`, `and`, `been`,
	`have`, `in`, `will`, `on`, `does`, `yourselves`, `then`, `that`, `because`, `what`, `over`, `why`, `so`,
	`can`, `did`, `not`, `now`, `under`, `he`, `you`, `herself`, `has`, `just`, `where`, `too`, `only`, `myself`,
	`which`, `those`, `i`, `after`, `few`, `whom`, `t`, `being`, `if`, `theirs`, `my`, `against`, `a`, `by`,
	`doing`, `it`, `how`, `further`, `was`, `here`, `than`}

type WordCount = struct {
	Word  string
	Count int
}

func CountWords(f *os.File) map[string]int {
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
			counts := CountWords(f)
			wcslice := make([]WordCount, 0, len(counts))
			for word, count := range counts {
				wcslice = append(wcslice, WordCount{word, count})
			}
			sort.Slice(wcslice, func(i, j int) bool {
				return wcslice[i].Count > wcslice[j].Count
			})
			fmt.Printf("%v\n", wcslice)
		}
	}
}
