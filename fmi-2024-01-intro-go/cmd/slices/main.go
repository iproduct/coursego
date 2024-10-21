package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("s1 = ", s1, len(s1), cap(s1))
	s2 := s1[3:8]
	fmt.Println("s2 = ", s2, len(s2), cap(s2))

	// Sieve of Eratosten
	n := 10000
	p := 2
	erat := make([]bool, n+1)
	sqrtn := int(math.Sqrt(float64(n)))
	p = 2
	for p <= sqrtn {
		for i := 2 * p; i < n+1; i += p {
			erat[i] = true
		}
		for {
			p++
			if !erat[p] {
				break
			}
		}
	}
	fmt.Printf("erat[:20] = %v\n", erat[:20])
	i := 0
	for p := 2; p < n+1; p++ {
		if !erat[p] {
			fmt.Printf("%5d ", p)
			i++
			if i%20 == 0 {
				fmt.Println()
			}
		}
	}

}
