// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple test: enumeration of type int starting at 0.

package main

import "fmt"

//go:generate stringer -type=Day
type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println()

	//Pills
	fmt.Println(Placebo)
	fmt.Println(Aspirin)
	fmt.Println(Ibuprofen)
	fmt.Println(Paracetamol)
	fmt.Println(Acetaminophen)

}

