package main

import "fmt"

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

func (slice ByteSlice) String() string { // implements fmt.Stringer
	return string(slice)
}

func test(val interface{}) string {
	switch val.(type) {
	case func() string:
		return val.(func() string)()
	case string:
		return val.(string)
	case fmt.Stringer:
		return val.(fmt.Stringer).String()
	default:
		return "Type not recognized"
	}
}

func main() {
	// 1) string
	fmt.Printf("test(\"abcd\"): %s\n", test("abcd"))

	// 2) Stringer
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days - ", 7)
	b.AppendPointer([]byte("APPENDED\n"))
	fmt.Printf("test(ByteSlice): %s\n", test(b))

	// 3) func() string
	f := func() string {
		return "Returned from function f: func() string"
	}
	fmt.Printf("test(func() string): %s\n", test(f))

}
