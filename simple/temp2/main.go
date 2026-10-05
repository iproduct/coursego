package main

import "fmt"

var val interface{}

func main() {
	//val := func() string {
	//	return "Returned from function f: func() string"
	//}
	switch str := val.(type) {
	case string:
		fmt.Print(str)
	case func() string:
		fmt.Print(str())
	case func():
		fmt.Print("Void function")
	case fmt.Stringer:
		fmt.Print(str.String())
	default:
		fmt.Print("Type not recognized")
	}
}
