package main

import "fmt"

// Append appends the contents of t to the end of s and returns the result.
// If s has enough capacity, it is extended in place; otherwise a
// new array is allocated and returned.
func Append[T any](s []T, t ...T) []T {
	lens := len(s)
	tot := lens + len(t)
	if tot < 0 {
		panic(any("Append: cap out of range"))
	}
	if tot > cap(s) {
		news := make([]T, tot, 2 * tot)
		copy(news, s)
		s = news
	}
	s = s[:tot]
	copy(s[lens:], t)
	return s
}

// Copy copies values from t to s, stopping when either slice is
// full, returning the number of values copied.
func Copy[T any](s, t []T) int {
	i := 0
	for ; i < len(s) && i < len(t); i++ {
		s[i] = t[i]
	}
	return i
}

func main() {
	s := Append([]int{1, 2, 3}, 4, 5, 6)
	fmt.Println(s) // Now s is []int{1, 2, 3, 4, 5, 6}.
	Copy(s[3:], []int{7, 8, 9})
	fmt.Println(s) // Now s is []int{1, 2, 3, 7, 8, 9}
}
