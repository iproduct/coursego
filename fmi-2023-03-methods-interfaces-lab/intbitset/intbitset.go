package intbitset

import (
	"bytes"
	"fmi-2023-03-methods-interfaces-lab/stringutil"
	"fmt"
	"strconv"
)

// IntBitset is an efficient implementation of set in small ints
type IntBitSet struct {
	data []uint64
}

// New creates empty IntBitSet
func New() *IntBitSet {
	return &IntBitSet{
		data: []uint64{},
	}
}

// BitString returns internal bit represention of a IntBitSet
func (s *IntBitSet) BitString() string {
	var buf bytes.Buffer
	buf.WriteString("[" + strconv.Itoa(len(s.data)) + "]{")
	for _, word := range s.data {
		buf.WriteString(stringutil.Reverse(fmt.Sprintf("%064b", word)))
	}
	buf.WriteString("}")
	return buf.String()
}

// Has chechs if an element belong to the set
func (s *IntBitSet) Has(elem int) bool {
	word, bit := elem/64, uint(elem%64)
	return word < len(s.data) && s.data[word]&(1<<bit) != 0
}

// Add adds an element to the set
func (s *IntBitSet) Add(elem int) {
	word, bit := elem/64, uint(elem%64)
	for word >= len(s.data) {
		s.data = append(s.data, 0)
	}
	s.data[word] |= 1 << bit
}

func (s *IntBitSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	for i, word := range s.data {
		if word == 0 {
			continue
		}
		for b := 0; b < 64; b++ {
			if word&(1<<b) != 0 {
				if buf.Len() > 1 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+b)
			}
		}
	}
	buf.WriteString("}")
	return buf.String()
}
