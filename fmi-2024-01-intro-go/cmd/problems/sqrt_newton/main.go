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
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
