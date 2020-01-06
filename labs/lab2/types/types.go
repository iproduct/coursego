package main

import (
	"fmt"
)

// ByteSlice models a slice of bytes
type ByteSlice []byte

//Append appends data and returns the new slice
func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
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
	fmt.Fprintf(&bs, "abcd")
	fmt.Printf("%T -> %#[1]v", bs)
}
