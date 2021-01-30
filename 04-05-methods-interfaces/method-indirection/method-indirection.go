package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// Pointer receiver methods can be called
	//with both values and pointers
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 5)

	p := &Vertex{4, 3}
	p.Scale(5)
	ScaleFunc(p, 2)
	fmt.Println(v, p)

	// Value receiver methods can be called
	//with both values and pointers
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
