package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(pi(5000))
	fmt.Printf("%#32.30f\n", piEpsilon(.0000001))
	fmt.Printf("%#32.30f\n", math.Pi)
}

// pi launches n goroutines to compute an
// approximation of pi.
func pi(n int) float64 {
	f := 0.0
	for k := 0; k <= n; k++ {
		f += term(float64(k))
	}
	return f
}

func piEpsilon(epsilon float64) float64 {
	f, newTerm, oldTerm := 0.0, 0.0, 0.0
	k := 0.0
outer:
	for {
		newTerm = term(k)
		f += newTerm
		if math.Abs(newTerm-oldTerm) < epsilon {
			break outer
		}
		oldTerm = newTerm
		k++
	}
	return f
}

func term(k float64) float64 {
	return 4 * math.Pow(-1, k) / (2*k + 1)
}
