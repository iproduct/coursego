package main

import (
	"fmt"
	"sort"
)

type Sequence []int // Copy method copies the current to a new Sequence
func (s Sequence) Copy() Sequence {
	result := make([]int, len(s))
	copy(result, s)
	return result
} // String method sorts elements and returns them as string
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

type Stringer interface {
	String() string
}

func String(value interface{}) string {
	//switch str := value.(type) { // type switch
	//case string:
	//	return str
	//case Stringer:
	//	return str.String()
	//default:
	//	return "Error: Not stringable"
	//}

	if str, ok := value.(string); ok { // The same as above with type assertion
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	} else {
		return "Error: Not stringable"
	}
}

func main() {
	data := []int{42, 23, 18, 95, 16}
	s1 := Sequence(data)
	//s2 := s1.Copy()
	//s2.String()
	//fmt.Printf("data = %v, s2 = %v, s2 = %v", data, s1, s2)

	var value interface{} // Value provided by caller.
	value = "abcd"
	fmt.Println(String(value))
	value = s1
	fmt.Println(String(value))
	value = 123
	fmt.Println(String(value))
}
