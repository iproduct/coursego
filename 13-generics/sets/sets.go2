package main

import "fmt"

// Set is a set of values.
type Set[T comparable] map[T]struct{}

// Make returns a set of some element type.
func Make[T comparable]() Set[T] {
	return make(Set[T])
}

// Add adds v to the set s.
// If v is already in s this has no effect.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Delete removes v from the set s.
// If v is not in s this has no effect.
func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Contains reports whether v is in s.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Len reports the number of elements in s.
func (s Set[T]) Len() int {
	return len(s)
}

// Iterate invokes f on each element of s.
// It's OK for f to call the Delete method.
func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

func main() {
	// Create a set of ints.
	// We pass int as a type argument.
	// Then we write () because Make does not take any non-type arguments.
	// We have to pass an explicit type argument to Make.
	// Function argument type inference doesn't work because the
	// type argument to Make is only used for a result parameter type.
	s := Make[int]()

	// Add the value 1 to the set s.
	s.Add(1)
	s.Add(5)
	s.Add(7)
	s.Add(11)

	// Check that s does not contain the value 2.
	if s.Contains(2) { panic("unexpected 2") }

	s.Iterate(func(n int){
		fmt.Println(n)
	})

}
