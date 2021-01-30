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
	number.n++
}

func main() {
	//type Named interface {
	//	Name() string
	//}
	//greeting := func (thing Named) string {
	//	return "Hello " + thing.Name()
	//}
	//greeting(nil)

	var n *myNumber
	(*n).plusOne()
	fmt.Println(n.n)
}
