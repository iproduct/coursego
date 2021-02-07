package main

import (
	"fmt"
	"math"
)

import "image/color"

// Vertex represents a 2D point
type Vertex struct {
	X, Y float64
}

// Distance calcualtes the distance to the other Vertice
func (v Vertex) Distance(other Vertex) float64 {
	return math.Hypot(other.X-v.X, other.Y-v.Y)
}

// Scale method scales the Vertex coordinates by a factor of f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type ColorVertex struct {
	Vertex
	Color color.RGBA
}

func main() {
	green := color.RGBA{0, 255, 0, 255}
	yellow := color.RGBA{255, 255, 0, 255}
	cv1 := ColorVertex{Vertex{2, 3}, green}
	cv2 := ColorVertex{Vertex{6, 6}, yellow}
	fmt.Println(cv1.Distance(cv2.Vertex)) // 5
	cv1.Scale(4)
	cv2.Scale(4)
	fmt.Println(cv1.Distance(cv2.Vertex))// 20
	// cv1.Distance(cv2) // no method cv1.Distance(ColorVertex)

	a := Vertex{2, 7}
	b := Vertex{5, 3}

	distance := Vertex.Distance  // method expression
	fmt.Println(distance(a, b))  // 5
	fmt.Printf("%T\n", distance) // func(intbitset_main.Vertex, intbitset_main.Vertex) float64

	scale := (*Vertex).Scale // method expression
	scale(&a, 2)
	fmt.Println(a)            // {4 14}
	fmt.Printf("%T\n", scale) // func(*intbitset_main.Vertex, float64)

	scaleB := (&b).Scale       // method value
	fmt.Printf("%T\n", scaleB) // func(float64)
	scaleB(2)
	fmt.Printf("Sacling b with factor 2: b now is %f\n", b) //{10 6}

	distanceFromA := a.Distance                                  // method value
	fmt.Printf("%T\n", distanceFromA)                            // func(*Vertex, float64)
	fmt.Printf("Distance from A of B is %f\n", distanceFromA(b)) //10

	type ColorVertexP struct {
		*Vertex
		Color color.RGBA
	}

	cvp1 := ColorVertexP{&Vertex{2, 7}, green}
	cvp2 := ColorVertexP{&Vertex{5, 3}, yellow}
	fmt.Println(cvp1.Distance(*cvp2.Vertex)) // "5"
	cvp1.Vertex = cvp2.Vertex
	cvp1.Scale(3)
	fmt.Println(*cvp1.Vertex, *cvp2.Vertex) // {15 9} {15 9}
}
