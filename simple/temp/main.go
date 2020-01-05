package main

import (
	"fmt"
	"sort"
)

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

type Stringer interface {
	String() string
}

var value interface{} // Value provided by caller.

func main() {
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}

	if str, ok := value.(string); ok {
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	}
}