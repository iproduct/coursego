package main

import "fmt"

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}
func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}
func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}
func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%v, %p\n", b, &b)
	b.Write([]byte{42})
	fmt.Printf("%v, %p\n", b, &b)
	bslice := make([]byte, 1, 1)
	bslice[0] = 108
	fmt.Printf("%v, %p\n", bslice, &bslice)
	bslice = append(bslice, 12, 50, 70)
	b2 := b.Append(bslice)
	fmt.Printf("%v, %p\n", bslice, &bslice)
	fmt.Printf("%v, %p\n", b2, &b2)

}
