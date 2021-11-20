package main

import "fmt"

//type Named interface {
//	Name() string
//}
//
//func greeting(thing Named) string {
//	return "Hello " + thing.Name()
//}

type myNumber struct {
	n int
}

func (number *myNumber) plusOne() {
	if number!=nil {
		number.n++
	}
}

func main() {
	type Named interface {
		Name() string
	}
	greeting := func(thing Named) string {
		if thing == nil {
			return "Hello Anonymous"
		}
		return "Hello " + thing.Name()
	}
	fmt.Println(greeting(nil))

	var n *myNumber
	n.plusOne()
	n.plusOne()
	n.plusOne()
	fmt.Println(n)
}
