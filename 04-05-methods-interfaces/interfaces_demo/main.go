package main

import (
	"fmt"
	"github.com/iproduct/coursego/04-05-methods-interfaces/interfaces"
	"math"
)

func main() {
	var point interfaces.Printer
	//point = interfaces.CartesianPoint{3, 4}
	point = interfaces.PolarPoint{5, math.Pi / 2}
	fmt.Printf("%T -> %+v", point, point)
	point.Print()
}
