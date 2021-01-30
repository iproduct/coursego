package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{5: "E", 1: "A", 7: "G", 2: "B", 6: "F", 3: "C", 4: "D"}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}
