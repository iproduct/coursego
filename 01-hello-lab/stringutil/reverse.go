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
	fmt.Printf("%#v, %#v\n", stringHeader, unsafe.Pointer(&s))
	r := []rune(s)
	runeSliceHeader := (*SliceHeader)(unsafe.Pointer(&r))
	fmt.Printf("%#v, %#v, r[0] -> %p\n", runeSliceHeader, unsafe.Pointer(&r), &r[0])

	fmt.Printf("%T\n", r)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	result := string(r)
	resultStringHeader := (*StringHeader)(unsafe.Pointer(&result))
	fmt.Printf("%#v, %p\n", resultStringHeader, unsafe.Pointer(&result))
	return result
}
