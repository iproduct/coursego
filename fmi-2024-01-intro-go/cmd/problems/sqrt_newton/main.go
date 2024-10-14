package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z, iter := 1.0, 0
	// TODO implement sqrt using newton formula z -= (z*z - x)/ (2 * z)
	return z, iter
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
