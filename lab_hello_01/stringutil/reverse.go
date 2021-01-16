package stringutil

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	fmt.Printf("%T\n", r)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
