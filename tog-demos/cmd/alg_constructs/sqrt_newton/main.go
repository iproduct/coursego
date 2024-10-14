package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	i := 0
	for {
		z -= (z*z - x) / (2 * z)
		i++
		if z*z-x < 1e-20 {
			break
		}
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
