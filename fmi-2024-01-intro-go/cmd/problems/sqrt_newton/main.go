package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z, iter, epsilon := 1.0, 0, 1e-10
	epsilon += 0 // just to compile
	// TODO implement sqrt using newton formula z -= (z*z - x)/ (2 * z)
	return z, iter
}

func main() {
	sqr, iter := Sqrt(2)
	fmt.Printf("%22.20f, %d iterations\n", sqr, iter)
	fmt.Printf("%22.20f\n", math.Sqrt(2))
}
