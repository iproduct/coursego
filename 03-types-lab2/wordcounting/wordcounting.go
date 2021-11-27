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
	word  string
	count int
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	splitRegex := regexp.MustCompile(`[\s,\.!?\"\'\\/;\[\]\(\)\d\{\}*:-]+`)
	stopSet := make(map[string]struct{}, len(stopwordsList))
	for _, sw := range stopwordsList {
		stopSet[sw] = struct{}{}
	}
	fmt.Println(stopSet)
	for _, filename := range files {
		err := processFile(filename, counts, splitRegex, stopSet)
		if err != nil {
			log.Printf("Error processing %s : %v\n", filename, err)
		}
	}
	words := make([]Entry, len(counts))
	i := 0
	for w, c := range counts {
		words[i] = Entry{w, c}
		i++
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].count > words[j].count
	})
	for _, e := range words[:20] {
		fmt.Printf("%-20.20s -> %5d\n", e.word, e.count)
	}

}

func processFile(filename string, counts map[string]int, splitRegex *regexp.Regexp, stopSet map[string]struct{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//words := strings.Fields(line)
		words := splitRegex.Split(line, -1)
		for _, w := range words {
			if _, ok := stopSet[strings.ToLower(w)]; !ok && len(w) > 1 {
				counts[w] += 1
			}
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}
