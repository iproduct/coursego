package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var stopwords = []string{`ourselves`, `hers`, `between`, `yourself`, `but`, `again`, `there`, `about`, `once`,
	`during`, `out`, `very`, `having`, `with`, `they`, `own`, `an`, `be`, `some`, `for`, `do`, `its`, `yours`,
	`such`, `into`, `of`, `most`, `itself`, `other`, `off`, `is`, `s`, `am`, `or`, `who`, `as`, `from`, `him`,
	`each`, `the`, `themselves`, `until`, `below`, `are`, `we`, `these`, `your`, `his`, `through`, `don`, `nor`,
	`me`, `were`, `her`, `more`, `himself`, `this`, `down`, `should`, `our`, `their`, `while`, `above`, `both`,
	`up`, `to`, `ours`, `had`, `she`, `all`, `no`, `when`, `at`, `any`, `before`, `them`, `same`, `and`, `been`,
	`have`, `in`, `will`, `on`, `does`, `yourselves`, `then`, `that`, `because`, `what`, `over`, `why`, `so`,
	`can`, `did`, `not`, `now`, `under`, `he`, `you`, `herself`, `has`, `just`, `where`, `too`, `only`, `myself`,
	`which`, `those`, `i`, `after`, `few`, `whom`, `t`, `being`, `if`, `theirs`, `my`, `against`, `a`, `by`,
	`doing`, `it`, `how`, `further`, `was`, `here`, `than`}

var stopwordsMap map[string]bool


func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		words := strings.Fields(input.Text())
		for _, word := range(words) {
			if !stopwordsMap[word] {
				counts[word]++
			}
		}
	}
}

func main() {
	stopwordsMap = make(map[string]bool)
	for _, word := range(stopwords) {
		stopwordsMap[word] = true
	}
	files := os.Args[1:]
	counts := make(map[string]int)

	if cont := len(files); cont != 0 {
		for _, fname := range files {
			fmt.Printf("Opening file: %s\n", fname)
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

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })

	for key, val := range counts {
		fmt.Printf("%-20.20s -> %5d\n", key, val)
	}
}
