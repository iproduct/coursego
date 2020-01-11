package main

import "fmt"

type Vertex struct{ X, Y int }

type Line struct{ A, B *Vertex }

var gv Vertex = Vertex{2, 5}
var gv2 Vertex = Vertex{12, 29}
var gl Line = Line{&gv, &gv2}

func test(l Line) {
	fmt.Printf("%v, same=%v\n", l, l.A == gl.A)
	l.B.X = 42
	fmt.Printf("%v, %v\n", *l.A, *l.B)
}

func main() {
	test(gl)
	fmt.Printf("%v, %v\n", *gl.A, *gl.B)
}
