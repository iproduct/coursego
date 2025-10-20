package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Copy method copies the current to a new Sequence
func (s Sequence) Copy() Sequence {
	result := make([]int, len(s))
	copy(result, s)
	return result
}

// String method sorts elements and returns them as string
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
