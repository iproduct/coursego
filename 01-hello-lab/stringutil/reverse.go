package stringutil

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data unsafe.Pointer
	Len  int
}

type SliceHeader struct {
	Data unsafe.Pointer
	Len  int
	Cap  int
}

func Reverse(s string) string {
	stringHeader := (*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%#v\n", stringHeader)
	r := []rune(s)
	runeSliceHeader := (*StringHeader)(unsafe.Pointer(&r))
	fmt.Printf("%#v\n", runeSliceHeader)

	fmt.Printf("%T\n", r)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	result := string(r)
	resultStringHeader := (*StringHeader)(unsafe.Pointer(&result))
	fmt.Printf("%#v\n", resultStringHeader)
	return result
}
