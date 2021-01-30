package main

import (
	"fmt"
)

// ByteSlice type
type ByteSlice []byte

// Append appends data and returns new bvyte slice
func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}

// AppendPointer appends data to existing ByteSlice
func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days - ", 7)
	//b = b.Append([]byte("APPENDED\n"))
	(*ByteSlice).AppendPointer(&b, []byte("APPENDED\n"))
	fmt.Printf("%#v", string(b))
}
