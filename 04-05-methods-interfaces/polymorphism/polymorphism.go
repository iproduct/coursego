package main

import (
	"fmt"
	"github.com/iproduct/coursego/04-05-methods-interfaces/interfaces"
	"math"
)

// PrintPoints prints all points given as argument
func PrintPoints(points []interfaces.Point) {
	for _, p := range points {
		p.Print()
		fmt.Printf(" -> (%f, %f)\n", p.XCoord(), p.YCoord())
	}
}

func main() {
	var p1, p2 interfaces.Point
	p1 = interfaces.CartesianPoint{X: 1, Y: 2}
	p2 = interfaces.PolarPoint{R: 5, A: math.Pi / 2}
	p1.Print()
	p2.Print()
	fmt.Printf("\np2 in cartesian coordinates: (%f, %f)\n\n", p2.XCoord(), p2.YCoord())

	points := []interfaces.Point{p1, p2}
	points = append(points,
		interfaces.PolarPoint{R: 5, A: math.Pi / 4},
		interfaces.PolarPoint{R: 5, A: math.Pi})
	PrintPoints(points)
}
