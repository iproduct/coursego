package main

import (
	"fmt"
)

// ByteSlice models a slice of bytes
type ByteSlice []byte

//Append appends data and returns the new slice
func (slice *ByteSlice) Append(data []byte) {
	*slice = append([]byte(*slice), data...)
}

//Append appends data and returns the new slice
func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	n = len(data)
	err = nil
	return
}

func main() {
	var bs ByteSlice
	bs.Append([]byte("WXYZ"))
	fmt.Fprintf(&bs, "abcdef")
	fmt.Printf("%T -> %#[1]v , string: %[1]s", bs)
}
