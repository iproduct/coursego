package main

import "fmt"
import "github.com/iproduct/coursego/fmi-2025-01-intro/byteslice"

func main() {
	var b byteslice.ByteSlice
	fmt.Fprintf(&b, "This week has %d days\n", 7)
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
