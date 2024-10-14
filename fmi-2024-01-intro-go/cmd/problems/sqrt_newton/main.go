package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (z float64, iter int) {
	z, iter, epsilon := 1.0, 0, 1e-15
	// TODO implement sqrt using newton formula z -= (z*z - x)/ (2 * z)
	//for math.Abs(z*z-x) > epsilon { // while loop
	//	z -= (z*z - x) / (2 * z)
	//	iter++
	//}
	for { // do-while loop
		z -= (z*z - x) / (2 * z)
		iter++
		if math.Abs(z*z-x) < epsilon {
			return
		}
	}
}

func main() {
	sqr, iter := Sqrt(2)
	fmt.Printf("%22.20f, %d iterations\n", sqr, iter)
	fmt.Printf("%22.20f\n", math.Sqrt(2))
}
