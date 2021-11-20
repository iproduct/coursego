package main

import (
	"fmt"
	"math"
)

// Vertex represents a 2D point
type Vertex struct {
	X, Y float64
}

// Distance calcualtes the distance to the other Vertice
func (v *Vertex) Distance(other Vertex) float64 {
	return math.Hypot(other.X-v.X, other.Y-v.Y)
}

// Scale method scales the Vertex coordinates by a factor of f
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Path represents a sequence of Vertices
// A nil Path represents empty sequence
type Path []Vertex

// Distance method calculates distance over the path
func (p *Path) Distance() (dist float64) {
	dist = 0
	if *p == nil || len(*p) == 0 {
		return 0
	}
	v1 := (*p)[0]
	var v2 Vertex
	for i := 1; i < len(*p); i++ {
		v2 = (*p)[i]
		dist += v1.Distance(v2)
		v1 = v2
	}
	return
}

// Scale method scales the Vertex coordinates by a factor of f
func (p *Path) Scale(f float64) {
	for  i := 0; i < len(*p); i ++ {
		(*p)[i].Scale(f)
	}
}

func main() {
	var path Path
	path = Path{{1, 1}, {4, 5}, {4, 1}, {1, 1}}
	(&path).Scale(10)
	fmt.Printf("Path: %#v\n", path)
	fmt.Println("Perimeter = ", path.Distance())
}
