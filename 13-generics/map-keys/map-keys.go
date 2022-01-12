package main

import "fmt"

// Keys returns the keys of the map m in a slice.
// The keys will be returned in an unpredictable order.
// This function has two type parameters, K and V.
// Map keys must be comparable, so key has the predeclared
// constraint comparable. Map values can be any type.
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Values[K comparable, V any](m map[K]V) []V {
	r := make([]V, 0, len(m))
	for k := range m {
		r = append(r, m[k])
	}
	return r
}

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
	// Now k is either []int{1, 2, 3, 4} or a permutation of it.
	fmt.Printf("Keys (%T): %[1]v\n", Keys(m))
	fmt.Printf("Values (%T): %[1]v\n", Values(m))
}
