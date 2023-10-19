package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("pi(5000)        = %v\n", piEpsilon(0.0000001))
	fmt.Printf("piEpsilon(5000) = %v\n", pi(5000))
	fmt.Printf("math.Pi         = %v\n", math.Pi)
}

func pi(n int) float64 {
	f := 0.0
	for k := 0; k <= n; k++ {
		f += term(k)
	}
	return f
}

func piEpsilon(epsilon float64) float64 {
	f := 0.0
	// your solution here ...
	return f
}

func term(n int) float64 {
	k := float64(n)
	return 4 * math.Pow(-1, k) / (2*k + 1)
}
