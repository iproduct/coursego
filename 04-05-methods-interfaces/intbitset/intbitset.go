package intbitset

import (
	"bytes"
	"fmt"
	"github.com/iproduct/coursego/04-05-methods-interfaces/stringutil"
	"strconv"
)

// IntBitSet efficient implementation of set in small ints
type IntBitSet struct {
	data []uint64
}

// BitString method returns internal bit representation as string
func (s *IntBitSet) bitString() string {
	var buf bytes.Buffer
	buf.WriteString("[" + strconv.Itoa(len(s.data)) + "]{")
	for _, word := range s.data {
		buf.WriteString(stringutil.Reverse(fmt.Sprintf("%064b", word)))
	}
	buf.WriteByte('}')
	return buf.String()
}

//// Has method returns if x is an element
//func (s *IntBitSet) Has(x int) bool {
//	word, bit := x/64, uint(x%64)
//	return word < len(s.data) && s.data[word]&(1<<bit) != 0
//}
//
//// Add adds element to Int Set
//func (s *IntBitSet) Add(x int) {
//	word, bit := x/64, uint(x%64)
//	for word >= len(s.data) {
//		s.data = append(s.data, 0)
//	}
//	s.data[word] |= 1 << bit
//}
//
//// BitString method returns internal bit representation as string
//func (s *IntBitSet) BitString() string {
//	var buf bytes.Buffer
//	buf.WriteString("[" + strconv.Itoa(len(s.data)) + "]{")
//	for _, word := range s.data {
//		buf.WriteString(stringutil.Reverse(fmt.Sprintf("%064b", word)))
//	}
//	buf.WriteByte('}')
//	return buf.String()
//}
//
//// String method return string representation of the bitset
//func (s *IntBitSet) String() string {
//	var buf bytes.Buffer
//	buf.WriteByte('{')
//	for i, word := range s.data {
//		if word == 0 {
//			continue
//		}
//		for b := 0; b < 64; b++ {
//			if word&(1<<b) != 0 {
//				if buf.Len() > 1 {
//					buf.WriteByte(' ')
//				}
//				fmt.Fprintf(&buf, "%d", 64*i+b)
//			}
//		}
//	}
//	buf.WriteByte('}')
//	return buf.String()
//}
