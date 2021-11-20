package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleVal(f float64) Vertex {
	v.X = v.X * f
	v.Y = v.Y * f
	return v
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Printf("After scaling by 10: %+v\n", v)
	v = Vertex.ScaleVal(v, 10)
	(*Vertex).Scale((&v), 10)
	fmt.Println(v.Abs())
	fmt.Println(Vertex.Abs(v))

	fmt.Println((&v).Abs())
	fmt.Println((*Vertex).Abs(&v))
}
