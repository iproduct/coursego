package main

import "fmt"

type Vertex struct {
	X, Y   int
	colors []int
}

var (
	v1 = Vertex{1, 2, []int{111, 222}}               // has type Vertex
	v2 = Vertex{X: 1, Y: 2, colors: []int{111, 222}} // Y:0 is implicit
	v3 = Vertex{}                                    // X:0 and Y:0
	p  = &Vertex{1, 2, nil}                          // has type *Vertex
)

func main() {
	//fmt.Println(v1 == v2) // Compile error: Invalid operation: v1 == v2 (operator == not defined on Vertex)
	fmt.Println(v1, p, v2, v3)
	p := v1
	p.X = 1e9
	fmt.Println(v1)
}
