package main

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte  { return append(slice, data...) }
func (slice *ByteSlice) AppendPointer(data []byte) { *slice = append(*slice, data...) }
func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append(*slice, data...)
	return len(data), nil
}
func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days", 7)
	fmt.Printf("%v -> %s", b, string(b))
	fmt.Println(b.Append([]byte{97}))
	fmt.Printf("%v -> %s", b, string(b))
	(&b).AppendPointer([]byte{97})
	fmt.Printf("%v -> %s", b, string(b))
}
