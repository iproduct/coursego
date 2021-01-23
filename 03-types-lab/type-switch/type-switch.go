package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"strconv"
)

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
	switch str := val.(type) {
	case func() string:
		return str()
	case string:
		return str
	case fmt.Stringer:
		return str.String()
	default:
		return "Type not recognized"
	}
}

// Testing function type conversion from func(int) string -> MyFunc implmenting error interface
type MyFunc func(int) string

func (f MyFunc) Error() string {
	return "Error computing the function"
}

func test2(val func(int) string, n int) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = MyFunc(val) // conversion to type MyFunc implementing error interface
		}
	}()
	return val(n), nil
}

func main() {
	// 1) string
	fmt.Printf("test(\"abcd\"): %s\n", test("abcd"))

	// 2) Stringer
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days - ", 7)
	b.AppendPointer([]byte("APPENDED"))
	fmt.Printf("test(ByteSlice): %s\n", test(b))

	// 3) func() string
	f := func() string {
		return "Returned from function f: func() string"
	}
	fmt.Printf("test(func() string): %s\n", test(f))

	// 4) Test function type conversion from func(int) string -> MyFunc
	f2 := func(n int) string {
		if n < 0 {
			panic("Agument should not be negative")
		}
		return strconv.Itoa(n)
	}

	result, err := test2(f2, 42)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\ntest2(func(int) string): %s\n", result)

	result, err = test2(f2, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\ntest2(func(int) string): %s\n", result)

}
