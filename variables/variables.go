package main

import (
	"fmt"
	"strconv"
)

var global int

func main() {
	global = 42
	var i int
	i = 5
	fmt.Printf("%[2]v, %[1]v\n", global, i)
	var cp1, cp2 complex128
	cp1 = 12 + 3i
	cp2 = 7 + 5i
	cp3 := &cp1
	fmt.Printf("%v\n", *cp3+cp2)
	intPointer := new(int)
	*intPointer++
	fmt.Printf("%v\n", *intPointer)
	genericChannel := make(chan interface{})
	go func() {
		var i int64
		for i = 0; i < 5; i++ {
			genericChannel <- "abcd_" + strconv.FormatInt(i, 10)
		}
		close(genericChannel)
	}()
	for ok := true; ok; {
		var result interface{}
		result, ok = <-genericChannel
		if ok {
			fmt.Printf("%v\n", result)
		} else {
			fmt.Printf("END")
		}
	}
}
