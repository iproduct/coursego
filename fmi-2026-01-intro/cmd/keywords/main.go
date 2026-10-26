package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"unicode"
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

type Entry struct {
	Word  string
	Count int
}

var stopwords map[string]struct{}

func init() {
	stopwords = map[string]struct{}{}
	for _, w := range stopwordsList {
		stopwords[w] = struct{}{}
	}
}

func processFile(filename string, counts map[string]int, splitRegexp *regexp.Regexp) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := splitRegexp.Split(line, -1)
		for _, word := range words {
			word = strings.ToLower(word)
			if _, isStopword := stopwords[word]; !isStopword &&
				len(word) >= 2 || (len(word) >= 1 && unicode.IsDigit(rune(word[0]))) {
				counts[word]++
			}
		}
	}
	for word, count := range counts {
		stem := word[:len(word)-1]
		if _, ok := counts[stem]; ok && word[len(word)-1] == 's' {
			counts[stem] += count
			delete(counts, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run wordcount.go FILE_1, [FILE_2], ..., [FILE_N]")
		os.Exit(1)
	}
	files := os.Args[1:]
	absBase, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	splitRexexp := regexp.MustCompile("\\W")
	counts := map[string]int{}
	for _, file := range files {
		filename := filepath.Join(absBase, file)
		err := processFile(filename, counts, splitRexexp)
		if err != nil {
			fmt.Printf("Error reading file '%s': %v", file, err)
		}
	}
	var entries []Entry
	for word, count := range counts {
		entries = append(entries, Entry{word, count})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count == entries[j].Count {
			return entries[i].Word < entries[j].Word
		} else {
			return entries[i].Count > entries[j].Count
		}
	})
	fmt.Println("Counts: %v\n", entries[:20])
}
