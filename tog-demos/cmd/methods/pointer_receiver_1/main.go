package main

import (
	"math"
)

type Scaler interface {
	Scale(f float64)
}
type Abser interface {
	Abs() float64
}
type ScalerAbser interface {
	Scale(f float64)
	Abs() float64
}

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

func main() {
	//var sc Scaler
	var abser Abser
	var scabser ScalerAbser
	v := Vertex{3, 4}
	abser = v
	abser.Abs()
	scabser = &v
	scabser.Scale(10)
	//fmt.Println(v.Abs())
}
