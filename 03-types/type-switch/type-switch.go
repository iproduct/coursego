package main

import (
	"fmt"
)

//type StringReturnigFunc func() string

func test(val interface{}) string {
	switch str := val.(type) {
	case string:
		return str
	case func() string:
		return str()
	case func():
		return "Void function"
	case fmt.Stringer:
		return str.String()
	default:
		return "Type not recognized"
	}
}

// Testing function type conversion from func(int) string -> MyFunc implmenting error interface
type MyFunc func(int) string

func (f MyFunc) Error() string {
	return "Error computing the function"
}

func test2(val func(int) string, n int) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			err = MyFunc(val) // conversion to type MyFunc implementing error interface
		}
	}()
	return val(n), nil
}

func main() {
	a1 := [...]int{1, 2}
	a2 := a1
	a2[0] = 3
	fmt.Printf("%v, %v, %t\n", a1, a2, &a1 == &a2)
}
