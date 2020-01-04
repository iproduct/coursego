package main

import (
	"fmt"
)

type empty interface{}
type example interface {
	sampleMethod()
}

func main() {
	value := 42
	var e empty = value
	var f float32
	f = float32(e.(int))
	fmt.Printf("%f\n", f)
	switch e.(type) {
	case int:
		fmt.Printf("int type: %d\n", e)
	default:
		fmt.Printf("Invalid type: %T\n", e)
	}
	// This will panic at runtime
	var example example = e.(example)
	fmt.Printf("%v", example)
}
