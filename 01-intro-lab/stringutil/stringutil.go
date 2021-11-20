// Package stringutil contains utility functions for proceesing strings
package stringutil

// Reverse returns its argument string reversed rune-wise
func Reverse(s string) string {
	r := []rune(s) // slice of runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
