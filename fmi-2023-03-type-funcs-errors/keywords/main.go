package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Counts = map[string]int

type WordCount struct {
	Word  string
	Count int
}

var stopWordsList = []string{`ourselves`, `hers`, `between`, `yourself`, `but`, `again`, `there`, `about`, `once`,
	`during`, `out`, `very`, `having`, `with`, `they`, `own`, `an`, `be`, `some`, `for`, `do`, `its`, `yours`,
	`such`, `into`, `of`, `most`, `itself`, `other`, `off`, `is`, `s`, `am`, `or`, `who`, `as`, `from`, `him`,
	`each`, `the`, `themselves`, `until`, `below`, `are`, `we`, `these`, `your`, `his`, `through`, `don`, `nor`,
	`me`, `were`, `her`, `more`, `himself`, `this`, `down`, `should`, `our`, `their`, `while`, `above`, `both`,
	`up`, `to`, `ours`, `had`, `she`, `all`, `no`, `when`, `at`, `any`, `before`, `them`, `same`, `and`, `been`,
	`have`, `in`, `will`, `on`, `does`, `yourselves`, `then`, `that`, `because`, `what`, `over`, `why`, `so`,
	`can`, `did`, `not`, `now`, `under`, `he`, `you`, `herself`, `has`, `just`, `where`, `too`, `only`, `myself`,
	`which`, `those`, `i`, `after`, `few`, `whom`, `t`, `being`, `if`, `theirs`, `my`, `against`, `a`, `by`,
	`doing`, `it`, `how`, `further`, `was`, `here`, `than`}

var splitRegexp = regexp.MustCompile(`\W+`)

var stopWordsMap = make(map[string]struct{}, len(stopWordsList))

func init() {
	for _, stopWord := range stopWordsList {
		stopWordsMap[stopWord] = struct{}{}
	}
}

func ExtractKeywords(file *os.File) (Counts, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := splitRegexp.Split(scanner.Text(), -1)
		for _, word := range words {
			word = strings.ToLower(word)
			if _, found := stopWordsMap[word]; found || len(word) < 2 {
				continue
			}
			counts[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return counts, nil
}

func ProcessFile(fname string) (Counts, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ExtractKeywords(file)
}

func TopKeywords(counts Counts, topN int) []WordCount {
	wordCounts := make([]WordCount, len(counts))
	for word, count := range counts {
		wordCounts = append(wordCounts, WordCount{word, count})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})
	return wordCounts[:topN]
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: keywords <filename_1>, ... <filename_N>")
	}
	files := os.Args[1:]
	for _, fname := range files {
		counts, err := ProcessFile(fname)
		if err != nil {
			log.Printf("Error opening file %s: %v\n", fname, err)
			continue
		}
		//fmt.Printf("%v\n", counts)
		fmt.Printf("%v\n", TopKeywords(counts, 15))
	}
}
