package main

import (
	"fmt"
	"github.com/iproduct/coursego/fmi-2025-01-intro/byteslice"
	"github.com/iproduct/coursego/fmi-2025-01-intro/sequence"
)

type Stringer interface {
	String() string
}

var value interface{}

// Value provided by caller.
func String(value interface{}) string {
	//switch str := value.(type) {
	//// type switch
	//case string:
	//	return str + " [string]"
	//case Stringer:
	//	return str.String() + " [Stringer]"
	//default:
	//	return "Unkown type"
	//}

	if str, ok := value.(string); ok { // The same as above with type assertion
		return str
	} else if str, ok := value.(Stringer); ok {
		return str.String()
	}
	return "Unkown type"
}

func main() {
	s := "Hello Golang"
	fmt.Println(String(s))
	seq := sequence.Sequence([]int{11, 22, 13, 54, 35, 16, 77, 8})
	fmt.Println(String(seq))
	byteslice := byteslice.New(10)
	byteslice.AppendPointer([]byte{1, 2, 3})
	fmt.Println(String(byteslice))
}
