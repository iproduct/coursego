package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
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

var stopwords map[string]struct{}

func init() {
	stopwords = make(map[string]struct{}, len(stopwordsList))
	for _, word := range stopwordsList {
		stopwords[word] = struct{}{}
	}
}

type Entry struct {
	Word  string
	Count int
}

func processFile(filename string, counts map[string]int, splitRegex *regexp.Regexp) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := splitRegex.Split(line, -1)
		for _, word := range words {
			match, err := regexp.MatchString("\\d+", word)
			if _, ok := stopwords[strings.ToLower(word)]; err != nil || match == true || len(word) < 2 || ok {
				continue
			}
			counts[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordcount file1, [flie2], ... [fileN]")
	}
	files := os.Args[1:]
	counts := make(map[string]int)
	for _, file := range files {
		err := processFile(file, counts, regexp.MustCompile("\\W"))
		if err != nil {
			log.Printf("file '%s' not found: %v\n", file, err)
		}
	}
	entries := make([]Entry, 0, len(counts))
	for word, count := range counts {
		entries = append(entries, Entry{word, count})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Count == entries[j].Count {
			return entries[i].Word > entries[j].Word
		} else {
			return entries[i].Count > entries[j].Count
		}

	})
	for _, entry := range entries[:Min(20, len(entries))] {
		fmt.Printf("%s: %d\n", entry.Word, entry.Count)
	}
}
